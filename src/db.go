package rce

import (
	"database/sql"
	"log"
)

type DB interface {
	Seed(string)
	GetBalance(Institution) map[Currency]int
	GetInstitution(string) Institution
}

type PostgresDB struct {
	SqlConnection *sql.DB
}

func NewPostgresDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rce_admin dbname=rce_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	return &PostgresDB{SqlConnection: db}
}

func (PostgresDB) GetBalance(institution Institution) map[Currency]int {
	return nil
}
