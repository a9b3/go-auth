package services

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// User is the definition of a user.
type User struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	Password string `json:"-"`
}

// UserGet returns user given the id.
func UserGet(db *sql.DB, id string) (User, error) {
	user := User{}
	err := db.QueryRow(`
	SELECT id, email, password FROM "user" WHERE id=$1
	`, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// UserCreate inserts new user row into the user table and returns the inserted
// id.
func UserCreate(db *sql.DB, email string, password string) (User, error) {
	user := User{}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	hashedPassword := string(hash)

	var id string
	err = db.QueryRow(`
	INSERT INTO "user" (email, password)
	VALUES ($1, $2)
	RETURNING id`, email, hashedPassword).Scan(&id)
	if err != nil {
		return user, err
	}

	user, err = UserGet(db, id)
	if err != nil {
		return user, err
	}

	return user, nil
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
