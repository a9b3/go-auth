package services

import (
	"fmt"
	"testing"

	// sql implementation
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	email := "cool"
	password := "asd"

	user, err := UserCreate(dbInstance, email, password)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(
		t,
		user.Email,
		email,
		fmt.Sprintf(
			`Created user's email "%s" must match original "%s"`,
			user.Email,
			email,
		),
	)
	assert.NotEqual(
		t,
		user.Password,
		password,
		fmt.Sprintf(
			`Created user's password "%s" must not equal original "%s"`,
			user.Password,
			password,
		),
	)

	assert.True(
		t,
		VerifyPassword(user.Password, password),
		fmt.Sprintf(
			`Created user's password "%s" must verify with original "%s"`,
			user.Password,
			password,
		),
	)

	_, err = UserCreate(dbInstance, email, password)
	assert.Error(t, err, `Creating user with same email should error`)
}
