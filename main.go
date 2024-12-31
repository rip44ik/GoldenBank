package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gitlab.com/xfx1/goldbank/api"
	db "gitlab.com/xfx1/goldbank/db/sqlc"
	"gitlab.com/xfx1/goldbank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot log config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
