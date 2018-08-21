package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthCheckResponse is the response format of HealthCheck
type HealthCheckResponse struct {
	DB bool
}

// CreateHealthCheckHandler acts as a closure for handler's dependencies.
func CreateHealthCheckHandler(dbClient *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbPingErr := dbClient.Ping()
		healthCheckResponse := HealthCheckResponse{
			DB: dbPingErr == nil,
		}

		healthCheckResponseJson, err := json.Marshal(healthCheckResponse)
		if err != nil {
			panic(fmt.Errorf("fatal error marshalling json: %s", err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(healthCheckResponseJson)
	}
}
