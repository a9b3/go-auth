package main

import (
	"net/http"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
	"github.com/esayemm/auth/handlers"
)

func main() {
	cfg := config.New(".env")
	err, dbInstance := db.DBClient(cfg)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health_check", handlers.CreateHealthCheckHandler(dbInstance))
	http.HandleFunc("/register", handlers.CreateRegisterHandler(dbInstance, cfg))
	http.HandleFunc("/authenticate", handlers.Authenticate)
	http.HandleFunc("/verify", handlers.Verify)
	http.HandleFunc("/logout", handlers.Logout)

	if err := http.ListenAndServe(":"+cfg["PORT"], nil); err != nil {
		panic(err)
	}
}
