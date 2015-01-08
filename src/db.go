package rcfe

import (
	"database/sql"
	"log"
)

type DB interface {
	Seed(string)
	GetBalance(Institution) map[Currency]int
	GetInstitution(string) Institution
}

func NewPostgresDB() *sql.DB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rcfe_admin dbname=rcfe_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return db
}
