package rce

import (
	"database/sql"
	"log"
)

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

func (PostgresDB) GetBalance(player Player) map[Currency]int {
	return nil
}

func (db PostgresDB) Seed(seed string) {
	_, err := db.Exec(seed)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}
