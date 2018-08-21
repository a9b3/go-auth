package handlers

import (
	"database/sql"
	"os"
	"testing"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
)

// Shared instance of db connection to be used by local tests
var dbClient *sql.DB
var cfg map[string]string

func TestMain(m *testing.M) {
	cfg = config.New("../.test.env")
	_, dbClient = db.DBClient(cfg)
	defer dbClient.Close()
	dbClient.Exec(`DELETE FROM "user";`)

	code := m.Run()
	dbClient.Exec(`DELETE FROM "user";`)
	os.Exit(code)
}
