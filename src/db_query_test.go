package rce_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func TestDB_CanBuildQuery(t *testing.T) {
	player := NewHedgeFund("hedge fund name")
	statement := CreatePlayerStatement(player)

	assert(t, strings.Index(statement, "insert into hedge_funds") != -1, fmt.Sprintf("should have found insert clause in '%s'", statement))
	assert(t, strings.Index(statement, "(name,dollars) values ('hedge fund name',100)") != -1, fmt.Sprintf("should have found values clause in '%s'", statement))
}

func TestDB_CanBuildUpdateQuery(t *testing.T) {
	NewTestDB()
	trade := NewTestTrade()
	statements := UpdateStatements(trade)
	traderCurrency := trade.Offer.TraderTransaction.Currency
	traderValuesClause := fmt.Sprintf("%s = %s + %d", traderCurrency, traderCurrency, trade.Offer.TraderTransaction.Amount)
	receiverCurrency := trade.Offer.ReceiverTransaction.Currency
	receiverValuesClause := fmt.Sprintf("%s = %s + %d", receiverCurrency, receiverCurrency, trade.Offer.ReceiverTransaction.Amount)

	assertEquals(t, 2, len(statements))

	traderUpdate := statements[0]
	receiverUpdate := statements[1]

	assert(t, strings.Index(traderUpdate, "update hedge_funds set ") != -1, fmt.Sprintf("should have found update clause in '%s'", traderUpdate))
	assert(t, strings.Index(traderUpdate, traderValuesClause) != -1, fmt.Sprintf("should have found values clause in '%s'", traderUpdate))
	assert(t, strings.Index(traderUpdate, "where name = ") != -1, fmt.Sprintf("should have found values clause in '%s'", traderUpdate))

	assert(t, strings.Index(receiverUpdate, "update banks set ") != -1, fmt.Sprintf("should have found update clause in '%s'", receiverUpdate))
	assert(t, strings.Index(receiverUpdate, receiverValuesClause) != -1, fmt.Sprintf("should have found values clause in '%s'", receiverUpdate))
	assert(t, strings.Index(receiverUpdate, "where name = ") != -1, fmt.Sprintf("should have found values clause in '%s'", receiverUpdate))
}
