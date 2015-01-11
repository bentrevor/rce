package rce

import (
	"fmt"
	"strings"
)

// TODO not sure if I even need this...
type SQLStatement interface {
	Statement(...interface{}) string
	Statements(...interface{}) []string
}

type Insert struct{}
type Update struct{}
type Create struct{}
type Drop struct{}
type Seed struct{}

func (Insert) Statement(player Player) string {
	tableName := player.TableName
	insertClause := fmt.Sprintf("insert into %s ", tableName)
	columnValues := getPlayerColumnValues(player)
	valuesClause := fmt.Sprintf("(%s) values (%s);", player.ColumnNames, columnValues)

	return insertClause + valuesClause
}

func (Seed) Statement(players []Player) string {
	statement := resetPlayersStatement()

	for _, player := range players {
		statement += Insert{}.Statement(player)
	}

	return statement
}

func (Update) Statements(trade Trade) []string {
	traderUpdate := updateBalanceStatement(trade.Offer.TraderTransaction, trade.Trader)
	receiverUpdate := updateBalanceStatement(trade.Offer.ReceiverTransaction, trade.Receiver)

	return []string{traderUpdate, receiverUpdate}
}

func resetPlayersStatement() string {
	statement := ""

	for _, tableName := range []string{"hedge_funds", "banks"} {
		statement += fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
		statement += fmt.Sprintf(`CREATE TABLE %s (
                                            id      SERIAL,
                                            name    VARCHAR(50) UNIQUE,
                                            dollars INTEGER
                                          );`, tableName)
	}

	return statement
}

func updateBalanceStatement(transaction Transaction, player Player) string {
	updateClause := fmt.Sprintf("update %s set ", player.TableName)
	currency := transaction.Currency
	valuesClause := fmt.Sprintf("%s = %s + %d ", currency, currency, transaction.Amount)
	whereClause := fmt.Sprintf("where name = '%s';", player.Name)

	return updateClause + valuesClause + whereClause
}

func getPlayerColumnValues(player Player) string {
	return fmt.Sprintf("'%s',%d", player.Name, player.Dollars)
}

func StringIncludes(str, substr string) bool {
	return strings.Index(str, substr) != -1
}
