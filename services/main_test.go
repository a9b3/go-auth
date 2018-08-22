package services

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
	"github.com/go-redis/redis"
)

// Shared instance of db connection to be used by local tests
var dbClient *sql.DB
var redisClient *redis.Client

func TestMain(m *testing.M) {
	cfg := config.New("../.test.env")

	dbClient, _ = db.DBClient(cfg)
	defer dbClient.Close()

	redisClient, _ = db.RedisClient(
		fmt.Sprintf(`%s:%s`, cfg["REDIS_HOST"], cfg["REDIS_PORT"]),
		cfg["REDIS_PASSWORD"],
		0,
	)
	defer redisClient.Close()

	dbClient.Exec(`DELETE FROM "user";`)
	redisClient.FlushAll()

	code := m.Run()
	dbClient.Exec(`DELETE FROM "user";`)
	redisClient.FlushAll()
	os.Exit(code)
}
