package rce_test

import (
	"database/sql"
	"log"

	. "github.com/bentrevor/rce/src"
)

func NewEmptyTestDB() *PostgresDB {
	// TODO not sure if ssl should be enabled in tests
	dbConfigs := "user=rce_admin dbname=rce_test sslmode=disable"
	db, err := sql.Open("postgres", dbConfigs)

	if err != nil {
		log.Fatal("\n\n\nfailure connecting to database with configs:\n  ", dbConfigs, "\nerr:\n  ", err, "\n\n\n")
	}

	return &PostgresDB{DB: db}
}
