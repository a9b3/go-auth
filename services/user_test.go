package services

import (
	"testing"

	// sql implementation
	_ "github.com/lib/pq"
)

func TestUserCreate(t *testing.T) {
	email := "cool"
	password := "asd"
	err, id := UserCreate(dbInstance, email, password)
	if err != nil {
		t.Fatalf(err.Error())
	}
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

	err, _ = UserCreate(dbInstance, email, password)
	if err == nil {
		t.Fatalf(`Creating user with same email should error`)
	}
}
