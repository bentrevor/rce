package rce

import (
	"database/sql"
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
