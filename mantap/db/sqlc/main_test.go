package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	/* Untuk image */
	dbSource = "postgres://root:mantap123@mantap_db:5432/mantap?sslmode=disable"
	/* untuk host */
	// dbSource = "postgres://root:mantap123@127.0.0.1:5432/mantap?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Tidak dapat konek ke database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
