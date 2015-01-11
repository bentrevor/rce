package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func TestDB_CanSeedFromSql(t *testing.T) {
	seed := `
	DROP TABLE IF EXISTS hedge_funds;
	CREATE TABLE hedge_funds (
		id      SERIAL,
		name    VARCHAR(50) UNIQUE,
		dollars INTEGER
	);

	INSERT INTO hedge_funds (name, dollars) VALUES ('test hedge fund', 1234);
`

	memoryDB := NewEmptyTestDB()
	memoryDB.Seed(seed)
	hedgeFund := NewHedgeFund("test hedge fund")

	assertEquals(t, 1234, memoryDB.GetBalance(hedgeFund)[Dollars])
}

func TestDB_CanGetBalance(t *testing.T) {
	memoryDB := NewEmptyTestDB()
	hedgeFund := NewHedgeFund("test hedge fund")
	seedStatement := Seed{}.Statement([]Player{hedgeFund})
	memoryDB.Seed(seedStatement)
	balance := memoryDB.GetBalance(hedgeFund)
	dollars := balance[Dollars]

	assertEquals(t, 100, dollars)
}
