package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func TestDB_CanGetBalance(t *testing.T) {
	memoryDB := NewTestDB()
	memoryDB.Seed(testSeed)
	hedgeFund := NewHedgeFund("test hedge fund")
	balance := memoryDB.GetBalance(hedgeFund)
	dollars := balance[Dollars]

	assertEquals(t, 100, dollars)
}

func TestDB_CanSeedFromSql(t *testing.T) {
	memoryDB := NewTestDB()
	memoryDB.Seed(testSeed)
	hedgeFund := NewHedgeFund("test hedge fund")

	balance := memoryDB.GetBalance(hedgeFund)
	dollars := balance[Dollars]

	assertEquals(t, 0, dollars)
}
