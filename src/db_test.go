package rce_test

import (
	"database/sql"
	"log"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

var seedStr = `
drop table if exists hedge_funds;
create table hedge_funds (
  id      serial,
  name    varchar(50),
  dollars integer,
  pesos   integer
);

insert into hedge_funds (name, dollars) values ('test hedge fund', 100);

drop table if exists banks;
create table banks (
  id      serial,
  name    varchar(50),
  dollars integer,
  pesos   integer
);

insert into banks (name, dollars) values ('test bank', 200);
`

func NewTestDB() *PostgresDB {
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rce_admin dbname=rce_test sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	postgresDB := &PostgresDB{DB: db}
	postgresDB.Seed(seedStr)
	return postgresDB
}

// func TestDB_CanGetBalance(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	balance := memoryDB.GetBalance(HedgeFund{Dollars: 10})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 10, dollars)
// }

// func TestDB_CanSeedFromPlayers(t *testing.T) {
// 	memoryDB := NewTestDB()
// 	memoryDB.Seed(nil)
// 	balance := memoryDB.GetBalance(HedgeFund{})
// 	dollars := balance[Dollars]

// 	assertEquals(t, 0, dollars)
// }
