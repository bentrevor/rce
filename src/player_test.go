package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"
)

var hedgeFund = NewHedgeFund("hedge fund name")
var bank = NewBank("bank name")

func TestPlayer_HasFunctionsForDBInfo(t *testing.T) {
	assertEquals(t, "banks", bank.TableName)
	assertEquals(t, "hedge_funds", hedgeFund.TableName)

	assertEquals(t, hedgeFund.ColumnNames, bank.ColumnNames)
}
