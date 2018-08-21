package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateHealthCheckHandler(dbClient, redisClient))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler must return 200 but got %v", status)
	}

	expected := `{"db":true,"redis":true}`
	if rr.Body.String() != expected {
		t.Errorf("response must be %v but got %v", expected, rr.Body.String())
	}
}
