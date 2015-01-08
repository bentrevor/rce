package rce_test

import (
	"database/sql"
	"log"
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func NewTestDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rce_admin dbname=rce_test sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{SqlConnection: db}
}

func TestDB_CanGetBalance(t *testing.T) {
	memoryDB := NewTestDB()

	balance := memoryDB.GetBalance(HedgeFund{})
	dollars := balance[Dollars]
	assertEquals(t, 0, dollars)
}
