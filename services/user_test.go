package services

import (
	"testing"

	// sql implementation
	_ "github.com/lib/pq"
)

func TestCreate(t *testing.T) {
	email := "cool"
	password := "asd"
	id := UserCreate(dbInstance, email, password)
	user := UserGet(dbInstance, id)

	if user.Email != email {
		t.Fatalf(`Created user's email "%s" must match original "%s"`, user.Email, email)
	}
	if user.ID != id {
		t.Fatalf(`Created user's id "%s" must match original "%s"`, user.ID, id)
	}
	if user.Password == password {
		t.Fatalf(`Created user's password "%s" must not equal original "%s"`, user.Password, password)
	}
	if !VerifyPassword(user.Password, password) {
		t.Fatalf(`Created user's password "%s" must verify with original "%s"`, user.Password, password)
	}
}
