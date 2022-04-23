package main

import (
	"database/sql"
	"github.com/tornvallalexander/go-backend-template/api"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"github.com/tornvallalexander/go-backend-template/utils"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not establish connection with db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("could not start server:", err)
	}
}
