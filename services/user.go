package services

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User is the definition of a user.
type User struct {
	Email    string
	ID       string
	Password string
}

// UserCreate inserts new user row into the user table.
func UserCreate(db *sql.DB, email string, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(fmt.Errorf("fatal error generating password: %s", err))
	}
	hashedPassword := string(hash)

	sqlStatement := `
	INSERT INTO "user" (email, password)
	VALUES ($1, $2)
	RETURNING id`
	var id string
	err = db.QueryRow(sqlStatement, email, hashedPassword).Scan(&id)
	if err != nil {
		panic(fmt.Errorf("fatal error inserting user: %s", err))
	}

	return id
}

// UserGet returns user given id
func UserGet(db *sql.DB, id string) User {
	sqlStatement := `
	SELECT id, email, password FROM "user" WHERE id=$1`
	user := User{}
	err := db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		panic(fmt.Errorf("fatal error querying user: %s", err))
	}

	return user
}

// VerifyPassword will return boolean for a password match
func VerifyPassword(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePass := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return false
	}

	return true
}
