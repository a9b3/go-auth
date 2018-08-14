package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateHealthCheckHandler(dbInstance))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler must return 200 but got %v", status)
	}

	expected := `{"DB":true}`
	if rr.Body.String() != expected {
		t.Errorf("response must be %v but got %v", expected, rr.Body.String())
	}
}
