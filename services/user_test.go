package services

import (
	"database/sql"
	"os"
	"testing"

	"github.com/esayemm/auth/config"
	"github.com/esayemm/auth/db"
	// sql implementation
	_ "github.com/lib/pq"
)

// Shared instance of db connection to be used by local tests
var dbInstance *sql.DB

func TestMain(m *testing.M) {
	cfg := config.New(".test.env")
	dbInstance = db.Open(cfg)

	// db.Exec(`DROP TABLE "user";`)

	code := m.Run()
	// db.Exec(`DROP TABLE "user";`)
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	UserCreate(dbInstance, "asd", "asd")
}
