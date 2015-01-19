package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"
)

func TestPlayer_HasFunctionsForDBInfo(t *testing.T) {
	var hedgeFund = NewHedgeFund("hedge fund name")
	var bank = NewBank("bank name")

	describe("player")
	it("has a name")
	assertEquals(t, "banks", bank.TableName)
	assertEquals(t, "hedge_funds", hedgeFund.TableName)

	specify("banks and hedge funds have the same db columns")
	assertEquals(t, hedgeFund.ColumnNames, bank.ColumnNames)
}
