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
	defer dbInstance.Close()
	dbInstance.Exec(`DELETE FROM "user";`)

	code := m.Run()
	dbInstance.Exec(`DELETE FROM "user";`)
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	id := UserCreate(dbInstance, "cool", "asd")
	user := UserGet(dbInstance, id)

	if user.Email != "cool" {
		t.Fatalf(`Created user's email "%s" must match original "%s"`, user.Email, "cool")
	}
	if user.ID != id {
		t.Fatalf(`Created user's id "%s" must match original "%s"`, user.ID, id)
	}
}
