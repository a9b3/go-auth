package services

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User is the definition of a user.
type User struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	Password string `json:"-"`
}

// UserCreate inserts new user row into the user table.
func UserCreate(db *sql.DB, email string, password string) (error, string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(fmt.Errorf("fatal error generating password: %s", err))
	}
	hashedPassword := string(hash)

	var id string
	err = db.QueryRow(`
	INSERT INTO "user" (email, password)
	VALUES ($1, $2)
	RETURNING id`, email, hashedPassword).Scan(&id)

	return err, id
}

// UserGet returns user given id
func UserGet(db *sql.DB, id string) User {
	user := User{}
	err := db.QueryRow(`
	SELECT id, email, password FROM "user" WHERE id=$1
	`, id).Scan(&user.ID, &user.Email, &user.Password)
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
