package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/esayemm/auth/services"
)

// RegisterBody is the required post body to register handler.
type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateRegisterHandler returns a handler.
func CreateRegisterHandler(db *sql.DB, cfg map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		registerBody := RegisterBody{}
		err = json.Unmarshal(jsn, &registerBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user, err := services.UserCreate(db, registerBody.Email, registerBody.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if cfg["SEND_EMAIL"] == "true" {
			err = services.Send(cfg["GMAIL_ACCOUNT"], cfg["GMAIL_PASSWORD"], user.Email, "Email Verification", "email")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	}
}
