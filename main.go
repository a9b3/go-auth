package main

import (
	"fmt"
	"net/http"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
	"github.com/esayemm/auth/handlers"
)

func main() {
	cfg := config.New(".env")

	dbInstance, err := db.DBClient(cfg)
	if err != nil {
		panic(err)
	}
	defer dbInstance.Close()

	redisClient, err := db.RedisClient(
		fmt.Sprintf(`%s:%s`, cfg["REDIS_HOST"], cfg["REDIS_PORT"]),
		cfg["REDIS_PASSWORD"],
		0,
	)
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	http.HandleFunc("/health_check", handlers.CreateHealthCheckHandler(dbInstance, redisClient))
	http.HandleFunc("/register", handlers.CreateRegisterHandler(dbInstance, redisClient, cfg))
	http.HandleFunc("/authenticate", handlers.Authenticate)
	http.HandleFunc("/verify", handlers.Verify)
	http.HandleFunc("/logout", handlers.Logout)

	if err := http.ListenAndServe(":"+cfg["PORT"], nil); err != nil {
		panic(err)
	}
}
