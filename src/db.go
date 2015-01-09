package rce

import (
	"database/sql"
	"fmt"
	"log"
)

type DbConnection interface {
	Seed([]Institution)
	GetBalance(Institution) map[Currency]int
	GetInstitution(string) Institution
}

type PostgresDB struct {
	*sql.DB
}

func NewPostgresDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rce_admin dbname=rce_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{DB: db}
}

func (PostgresDB) GetBalance(institution Institution) map[Currency]int {
	return nil
}

func CreateInstitutionStatement(institution Institution) string {
	tableName := institution.TableName()
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

func getColumnValues(institution Institution) string {
	return fmt.Sprintf("'%s',%d", institution.GetName(), institution.GetDollars())
}
