package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/esayemm/auth/services"
	"github.com/go-redis/redis"
)

// RegisterBody is the required post body to register handler.
type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateRegisterHandler returns a handler.
func CreateRegisterHandler(db *sql.DB, redisClient *redis.Client, cfg map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		registerBody := RegisterBody{}
		err = json.Unmarshal(jsn, &registerBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := services.UserCreate(db, registerBody.Email, registerBody.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if cfg["SEND_EMAIL"] == "true" {
			token, err := services.CreateVerificationToken(redisClient, user.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = services.Send(
				cfg["GMAIL_ACCOUNT"],
				cfg["GMAIL_PASSWORD"],
				user.Email,
				"Verification Token",
				fmt.Sprintf(`Your verification token is %s`, token),
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	}
}
