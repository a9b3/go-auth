package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

// HealthCheckResponse is the response format of HealthCheck
type HealthCheckResponse struct {
	DB    bool `json:"db"`
	Redis bool `json:"redis"`
}

// CreateHealthCheckHandler acts as a closure for handler's dependencies.
func CreateHealthCheckHandler(dbClient *sql.DB, redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbPingErr := dbClient.Ping()
		_, redisErr := redisClient.Ping().Result()

		healthCheckResponse := HealthCheckResponse{
			DB:    dbPingErr == nil,
			Redis: redisErr == nil,
		}

		healthCheckResponseJSON, err := json.Marshal(healthCheckResponse)
		if err != nil {
			panic(fmt.Errorf("fatal error marshalling json: %s", err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(healthCheckResponseJSON)
	}
}
