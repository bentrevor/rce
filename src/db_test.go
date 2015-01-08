package rcfe_test

import (
	"testing"

	. "github.com/bentrevor/rcfe/src"
)

type MemoryDB struct {
	institutions []Institution
}

func (MemoryDB) GetBalance(institution Institution) map[Currency]int {
	return map[Currency]int{
		Dollars: 0,
	}
}

func (db MemoryDB) Seed(s string) {
	hedgeFund := HedgeFund{Name: "hedge fund name", Dollars: 100}
	bank := Bank{Name: "bank name", Dollars: 200}

	db.institutions = []Institution{hedgeFund, bank}
}

func (db MemoryDB) GetInstitution(name string) Institution {
	for _, institution := range db.institutions {
		if institution.GetName() == name {
			return institution
		}
	}

	return nil
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}

func TestDB_CanGetBalance(t *testing.T) {
	memoryDB := NewMemoryDB()

	balance := memoryDB.GetBalance(HedgeFund{})
	dollars := balance[Dollars]
	assertEquals(t, 0, dollars)
}
