package rce

import "fmt"

func CreatePlayerStatement(player Player) string {
	tableName := player.TableName
	insertClause := fmt.Sprintf("insert into %s ", tableName)

	columnNames := getPlayerColumnNames()
	columnValues := getPlayerColumnValues(player)

	valuesClause := fmt.Sprintf("(%s) values (%s);", columnNames, columnValues)

	return insertClause + valuesClause
}

// TODO this is bad, if I change these columns I have to update it here too
func getPlayerColumnNames() string {
	return "name,dollars"
}

func getPlayerColumnValues(player Player) string {
	return fmt.Sprintf("'%s',%d", player.Name, player.Dollars)
}
