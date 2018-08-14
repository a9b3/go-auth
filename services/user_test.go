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
	email := "cool"
	password := "asd"
	id := UserCreate(dbInstance, email, password)
	user := UserGet(dbInstance, id)

	if user.Email != email {
		t.Fatalf(`Created user's email "%s" must match original "%s"`, user.Email, email)
	}
	if user.ID != id {
		t.Fatalf(`Created user's id "%s" must match original "%s"`, user.ID, id)
	}
	if user.Password == password {
		t.Fatalf(`Created user's password "%s" must not equal original "%s"`, user.Password, password)
	}
	if !VerifyPassword(user.Password, password) {
		t.Fatalf(`Created user's password "%s" must verify with original "%s"`, user.Password, password)
	}
}
