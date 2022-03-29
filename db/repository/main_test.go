package db

import (
	"database/sql"
	"github.com/Adetunjii/simplebank/util"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"testing"
)



var testRepository *Repository
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("An error occurred", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the database::: ", err)
	}

	testRepository = CreateNew(testDB)
	os.Exit(m.Run())
}
