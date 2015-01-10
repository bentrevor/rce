package rce

import (
	"database/sql"
	"log"
)

type DbConnection interface {
	Seed([]Player)
	GetBalance(Player) map[Currency]int
	GetPlayer(string) Player
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

func (PostgresDB) GetBalance(player Player) map[Currency]int {
	return nil
}
