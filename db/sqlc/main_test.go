package db

import (
	"IMChat/core/config"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {

	conf, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("load config err: ", err)
	}

	connPool, err := pgxpool.New(context.Background(), conf.Postgres.Source)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
