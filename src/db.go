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

func (PostgresDB) GetBalance(player Player) map[Currency]int {
	return nil
}

func seedDB(db *PostgresDB) {
	fmt.Println("seeding db...")

	_, err := db.Query(`
drop table if exists hedge_funds;
create table hedge_funds (
  id      serial,
  name    varchar(50),
  dollars integer,
  pesos   integer
);

insert into hedge_funds (name, dollars, pesos) values ('hf1', 1234, 123),
                                                      ('hf2', 5678, 567),
                                                      ('hf3', 9090, 909)
;

drop table if exists banks;
create table banks (
  id      serial,
  name    varchar(50),
  dollars integer,
  pesos   integer
);

insert into banks (name, dollars, pesos) values ('b1', 11111, 1111),
                                                ('b2', 22222, 2222),
                                                ('b3', 33333, 3333)
;
`)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}
