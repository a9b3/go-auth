package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esayemm/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	email := "foo"
	password := "123"

	jsonBytes, err := json.Marshal(RegisterBody{Email: email, Password: password})
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(CreateRegisterHandler(dbClient, cfg)).ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, `handler must return 200`)

	var user services.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, user.Email, email, `Created email must equal original`)
	assert.Empty(t, user.Password, "returned object must not contain password")
}
