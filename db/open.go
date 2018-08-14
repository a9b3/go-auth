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
		cfg["POSTGRESHOST"],
		cfg["POSTGRESPORT"],
		cfg["POSTGRESUSER"],
		cfg["POSTGRESPASSWORD"],
		cfg["POSTGRESDB"],
		cfg["POSTGRESSSLMODE"]),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
