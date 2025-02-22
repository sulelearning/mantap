package main

import (
	"database/sql"
	"log"

	"github.com/Zulhaidir/microservice/mantap/api"
	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	"github.com/Zulhaidir/microservice/mantap/util"
	_ "github.com/lib/pq"
)

func main() {
	/* Load Config */
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Tidak dapat konek ke database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Tidak dapat menjalankan server:", err)
	}
}
