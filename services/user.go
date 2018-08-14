package services

import (
	"database/sql"
	"fmt"
)

// User is the definition of a user.
type User struct {
	Name string
}

// Create inserts new user row into the user table.
func UserCreate(db *sql.DB, email string, password string) {
	sqlStatement := fmt.Sprintf(`
	INSERT INTO users (email, password)
	VALUES (%s, %s);`, email, password)
	fmt.Printf(sqlStatement)
}
