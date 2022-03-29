package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://postgres:secret@localhost:5432/simplebank?sslmode=disable"
)


var testRepository *Repository
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the database::: ", err)
	}

	testRepository = CreateNew(testDB)
	os.Exit(m.Run())
}
