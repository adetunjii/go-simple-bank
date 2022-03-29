package main

import (
	"database/sql"
	"github.com/Adetunjii/simplebank/api"
	db "github.com/Adetunjii/simplebank/db/repository"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"

)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://teej4y:password@localhost:5432/simplebank?sslmode=disable"
	address = "0.0.0.0:8080"
)


func main() {
	connection, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the database::: ", err)
	}

	store := db.CreateNewStore(connection)
	server := api.CreateNewServer(store)

	err = server.StartServer(address)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}