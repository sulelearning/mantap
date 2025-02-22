package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Zulhaidir/microservice/mantap/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Tidak dapat Load Config di TestMain:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Tidak dapat konek ke database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
