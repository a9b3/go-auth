package handlers

import (
	"database/sql"
	"os"
	"testing"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
)

// Shared instance of db connection to be used by local tests
var dbInstance *sql.DB

func TestMain(m *testing.M) {
	cfg := config.New("../.test.env")
	dbInstance = db.Open(cfg)
	defer dbInstance.Close()
	dbInstance.Exec(`DELETE FROM "user";`)

	code := m.Run()
	dbInstance.Exec(`DELETE FROM "user";`)
	os.Exit(code)
}
