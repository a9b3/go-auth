package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esayemm/auth/services"
)

func TestRegisterHandler(t *testing.T) {
	jsonBytes, err := json.Marshal(RegisterBody{Email: "foo", Password: "123"})
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRegisterHandler(dbInstance))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler must return 200 but got %v", status)
	}

	user := services.User{}
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Errorf(err.Error())
	}

	if user.Email != "foo" {
		t.Fatalf(`Created email must equal original`)
	}
	if user.Password != "" {
		t.Fatalf(`Returned object must not contain password`)
	}
}
