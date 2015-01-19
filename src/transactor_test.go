package rce_test

import (
	. "github.com/bentrevor/rce/src"

	"testing"

	_ "github.com/lib/pq"
)

func TestTransactor_CanExecuteATransaction(t *testing.T) {
	transactor := Transactor{}
	db := NewEmptyTestDB()
	hedgeFund := NewHedgeFund("test hedge fund")
	bank := NewBank("test bank")
	players := []Player{hedgeFund, bank}
	seedStatement := Seed{}.Statement(players)
	db.Seed(seedStatement)

	transaction := Transaction{
		Trader:   hedgeFund,
		Receiver: bank,
		Amount:   10,
		Currency: Dollars,
	}

	transactor.Execute(transaction, db)

	// assertEquals(t, 95, db.GetBalance(hedgeFund)[Dollars])
	// assertEquals(t, 205, db.GetBalance(bank)[Dollars])
}
