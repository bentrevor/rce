package rce

import (
	"fmt"
	"strings"
)

func CreatePlayerStatement(player Player) string {
	tableName := player.TableName
	insertClause := fmt.Sprintf("insert into %s ", tableName)

	columnValues := getPlayerColumnValues(player)

	valuesClause := fmt.Sprintf("(%s) values (%s);", player.ColumnNames, columnValues)

	return insertClause + valuesClause
}

func SeedStatement(players []Player) string {
	statement := resetTableStatement("hedge_funds")

	for _, player := range players {
		statement += CreatePlayerStatement(player)
	}

	return statement
}

func resetTableStatement(tableName string) string {
	statement := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)

	statement += `
CREATE TABLE hedge_funds (
  id      SERIAL,
  name    VARCHAR(50) UNIQUE,
  dollars INTEGER
);`

	return statement
}

func UpdateStatements(trade Trade) []string {
	traderUpdate := updateBalanceStatement(trade.Offer.TraderTransaction, trade.Trader)
	receiverUpdate := updateBalanceStatement(trade.Offer.ReceiverTransaction, trade.Receiver)

	return []string{traderUpdate, receiverUpdate}
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
