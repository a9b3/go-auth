package db

import (
	"database/sql"
	"fmt"
	// sql implementation
	_ "github.com/lib/pq"
)

// Open returns a *sql.DB
func DBClient(cfg map[string]string) (error, *sql.DB) {
	var db *sql.DB
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
		return err, db
	}
	if err := db.Ping(); err != nil {
		return err, db
	}

	return nil, db
}
