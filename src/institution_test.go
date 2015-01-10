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

func TestInstitution_HasFunctionsForDBInfo(t *testing.T) {
	assertEquals(t, "banks", bank.TableName())
	assertEquals(t, "hedge_funds", hedgeFund.TableName())

	assertEquals(t, InstitutionColumns, bank.ColumnNames())
	assertEquals(t, InstitutionColumns, hedgeFund.ColumnNames())
}
