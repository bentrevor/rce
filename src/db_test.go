package rce_test

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

	return &PostgresDB{DB: db}
}

// func TestDB_CanGetBalance(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	balance := memoryDB.GetBalance(HedgeFund{Dollars: 10})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 10, dollars)
// }

func TestDB_CanBuildQuery(t *testing.T) {
	institution := HedgeFund{Name: "hedge fund name", Dollars: 100}
	statement := CreateInstitutionStatement(institution)

	assert(t, strings.Index(statement, "insert into hedge_funds") != -1, fmt.Sprintf("should have found insert clause in '%s'", statement))
	assert(t, strings.Index(statement, "(name,dollars) values ('hedge fund name',100)") != -1, fmt.Sprintf("should have found values clause in '%s'", statement))
}

// func TestDB_CanSeedFromInstitutions(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	memoryDB.Seed(nil)
// 	balance := memoryDB.GetBalance(HedgeFund{})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 0, dollars)
// }
