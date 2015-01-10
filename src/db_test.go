package rce_test

import (
	"database/sql"
	"log"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func NewTestDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rce_admin dbname=rce_test sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{DB: db}
}

// func TestDB_CanGetBalance(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	balance := memoryDB.GetBalance(HedgeFund{Dollars: 10})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 10, dollars)
// }

// func TestDB_CanSeedFromPlayers(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	memoryDB.Seed(nil)
// 	balance := memoryDB.GetBalance(HedgeFund{})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 0, dollars)
// }
