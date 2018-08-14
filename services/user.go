package services

import (
	"database/sql"
	"fmt"
)

// User is the definition of a user.
type User struct {
	Email string
	ID    string
}

// UserCreate inserts new user row into the user table.
func UserCreate(db *sql.DB, email string, password string) string {
	sqlStatement := `
	INSERT INTO "user" (email, password)
	VALUES ($1, $2)
	RETURNING id`
	var id string
	err := db.QueryRow(sqlStatement, email, password).Scan(&id)
	if err != nil {
		panic(fmt.Errorf("fatal error inserting user: %s", err))
	}

	return id
}

// UserGet returns user given id
func UserGet(db *sql.DB, id string) User {
	sqlStatement := `
	SELECT id, email FROM "user" WHERE id=$1`
	user := User{}
	err := db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Email)
	if err != nil {
		panic(fmt.Errorf("fatal error querying user: %s", err))
	}

	return user
}
