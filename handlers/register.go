package handlers

import (
	"database/sql"
	"net/http"
)

// CreateRegisterHandler returns a handler.
func CreateRegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
