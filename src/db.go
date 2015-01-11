package rce

import (
	"database/sql"
	"fmt"
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

func (db PostgresDB) GetBalance(player Player) map[Currency]int {
	balance := make(map[Currency]int)

	query := fmt.Sprintf("select %s from %s;", player.ColumnNames, player.TableName)
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		log.Fatal("got an error: ", err)
	}

	for rows.Next() {
		var name string
		var dollars int
		rows.Scan(&name, &dollars)
		balance[Dollars] = dollars
	}

	fmt.Println(player.Name)
	return balance
}

func (db PostgresDB) Seed(seed string) {
	_, err := db.Exec(seed)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}
