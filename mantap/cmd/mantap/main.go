package main

import (
	"database/sql"
	"log"

	"github.com/Zulhaidir/microservice/mantap/api"
	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	/* Untuk image */
	dbSource = "postgres://root:mantap123@mantap_db:5432/mantap?sslmode=disable"
	/* untuk host */
	// dbSource = "postgres://root:mantap123@127.0.0.1:5432/mantap?sslmode=disable"

	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Tidak dapat konek ke database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Tidak dapat menjalankan server:", err)
	}
}
