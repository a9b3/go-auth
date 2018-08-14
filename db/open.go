package db

import (
	"database/sql"
	"fmt"
	// sql implementation
	_ "github.com/lib/pq"
)

// Open returns a *sql.DB
func Open(cfg map[string]string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg["POSTGRES_HOST"],
		cfg["POSTGRES_PORT"],
		cfg["POSTGRES_USER"],
		cfg["POSTGRES_PASSWORD"],
		cfg["POSTGRES_DB"],
		cfg["POSTGRES_SSLMODE"]),
	)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
