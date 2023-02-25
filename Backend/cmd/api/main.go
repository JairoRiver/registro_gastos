package main

import (
	"database/sql"
	"log"

	"github.com/JairoRiver/registro_gastos/tree/main/Backend/internal/api"
	db "github.com/JairoRiver/registro_gastos/tree/main/Backend/internal/db/sqlc"
	"github.com/JairoRiver/registro_gastos/tree/main/Backend/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
