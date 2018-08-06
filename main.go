package main

import (
	"fmt"
	"net/http"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/handlers"
)

func main() {
	cfg, _ := config.New()

	fmt.Printf("%+v", cfg)

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/authenticate", handlers.Authenticate)
	http.HandleFunc("/verify", handlers.Verify)
	http.HandleFunc("/logout", handlers.Logout)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		panic(err)
	}
}
