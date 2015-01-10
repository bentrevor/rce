package rce

import "fmt"

func CreateInstitutionStatement(institution Institution) string {
	tableName := institution.TableName
	insertClause := fmt.Sprintf("insert into %s ", tableName)

	columnNames := getInstitutionColumnNames()
	columnValues := getInstitutionColumnValues(institution)

	valuesClause := fmt.Sprintf("(%s) values (%s);", columnNames, columnValues)

	return insertClause + valuesClause
}

// TODO this is bad, if I change these columns I have to update it here too
func getInstitutionColumnNames() string {
	return "name,dollars"
}

func getInstitutionColumnValues(institution Institution) string {
	return fmt.Sprintf("'%s',%d", institution.Name, institution.Dollars)
}
