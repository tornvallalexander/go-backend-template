package db

import (
	"database/sql"
	"github.com/tornvallalexander/go-backend-template/utils"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("could not load environment variables:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
