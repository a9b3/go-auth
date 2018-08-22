package services

import (
	"testing"

	// sql implementation
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateVerificationToken(t *testing.T) {
	id := "123"

	token, err := CreateVerificationToken(redisClient, id)
	if err != nil {
		t.Fatal(err.Error())
	}

	allKeys, _ := redisClient.Keys("*").Result()

	assert.NotEmpty(t, token, "Token should not be empty")
	assert.Equal(t, len(allKeys), 1, "All keys should only have length of 1")
}

func TestGetVerificationToken(t *testing.T) {
	id := "123"
	token, _ := CreateVerificationToken(redisClient, id)

	val, err := GetVerificationToken(redisClient, id)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, token, val, "Created token should be gotten value")
}
