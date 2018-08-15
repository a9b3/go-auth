package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
func CreateRegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(fmt.Errorf("fatal error parsing request body: %s", err))
		}

		registerBody := RegisterBody{}
		err = json.Unmarshal(jsn, &registerBody)
		if err != nil {
			panic(fmt.Errorf("fatal error decoding json: %s", err))
		}

		err, id := services.UserCreate(db, registerBody.Email, registerBody.Password)
		if err != nil {
			panic(fmt.Errorf("fatal error creating user: %s", err))
		}

		user := services.UserGet(db, id)

		userJSON, err := json.Marshal(user)
		if err != nil {
			panic(fmt.Errorf("fatal error : %s", err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	}
}
