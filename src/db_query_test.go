package rce_test

import (
	"fmt"
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func TestDB_CanBuildSeedStatement(t *testing.T) {
	player := NewHedgeFund("hedge fund name")
	players := []Player{player}
	statement := SeedStatement(players)

	assert(t, StringIncludes(statement, "DROP TABLE IF EXISTS hedge_funds;"), "should have found drop table clause in statement")
	assert(t, StringIncludes(statement, "CREATE TABLE hedge_funds"), "should have found create table clause in statement")
	assert(t, StringIncludes(statement, CreatePlayerStatement(player)), "should have found create player clause in statement")
}

func TestDB_CanBuildInsertStatement(t *testing.T) {
	player := NewHedgeFund("hedge fund name")
	statement := CreatePlayerStatement(player)

	assert(t, StringIncludes(statement, "insert into hedge_funds"), "should have found insert clause")
	assert(t, StringIncludes(statement, "(name,dollars) values ('hedge fund name',100)"), "should have found values clause")
}

func TestDB_CanBuildUpdateStatement(t *testing.T) {
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

	assert(t, StringIncludes(traderUpdate, "update hedge_funds set "), "should have found update clause in '%s'")
	assert(t, StringIncludes(traderUpdate, traderValuesClause), "should have found values clause in '%s'")
	assert(t, StringIncludes(traderUpdate, "where name = "), "should have found values clause in '%s'")

	assert(t, StringIncludes(receiverUpdate, "update banks set "), "should have found update clause in '%s'")
	assert(t, StringIncludes(receiverUpdate, receiverValuesClause), "should have found values clause in '%s'")
	assert(t, StringIncludes(receiverUpdate, "where name = "), "should have found values clause in '%s'")
}
