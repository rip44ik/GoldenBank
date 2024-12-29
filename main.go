package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gitlab.com/xfx1/goldbank/api"
	db "gitlab.com/xfx1/goldbank/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:aswedD4321@localhost:5432/gold_bank?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
