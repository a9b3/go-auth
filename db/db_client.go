package db

import (
	"database/sql"
	"fmt"
	// sql implementation
	_ "github.com/lib/pq"
)

// DBClient returns a *sql.DB
func DBClient(cfg map[string]string) (*sql.DB, error) {
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
		return db, err
	}
	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
