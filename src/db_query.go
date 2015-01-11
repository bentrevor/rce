package rce

import "fmt"

func CreatePlayerStatement(player Player) string {
	tableName := player.TableName
	insertClause := fmt.Sprintf("insert into %s ", tableName)

	columnValues := getPlayerColumnValues(player)

	valuesClause := fmt.Sprintf("(%s) values (%s);", player.ColumnNames, columnValues)

	return insertClause + valuesClause
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
