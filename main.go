package main

import (
	"database/sql"
	"github.com/Adetunjii/simplebank/api"
	db "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/util"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"

)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	connection, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the database::: ", err)
	}

	store := db.CreateNewStore(connection)
	server := api.CreateNewServer(store)

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}