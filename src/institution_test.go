package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"
)

var hedgeFund = HedgeFund{
	Name:    "hedge fund name",
	Dollars: 50,
}

var bank = Bank{
	Name:    "bank name",
	Dollars: 100,
}

func TestHedgeFund_HasATableInTheDB(t *testing.T) {
	assertEquals(t, "hedge fund name", hedgeFund.Name)
	assertEquals(t, 50, hedgeFund.Dollars)
	assertEquals(t, "hedge_funds", hedgeFund.TableName())
}

func TestBank_HasATableInTheDB(t *testing.T) {
	assertEquals(t, "bank name", bank.Name)
	assertEquals(t, 100, bank.Dollars)
	assertEquals(t, "banks", bank.TableName())
}
