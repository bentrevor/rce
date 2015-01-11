package rce_test

import (
	"database/sql"
	"log"

	. "github.com/bentrevor/rce/src"
)

var testSeed = `
DROP TABLE IF EXISTS hedge_funds;
CREATE TABLE hedge_funds (
  id      SERIAL,
  name    VARCHAR(50) UNIQUE,
  dollars INTEGER
);

INSERT INTO hedge_funds (name, dollars) VALUES ('test hedge fund', 100);

DROP TABLE IF EXISTS banks;
CREATE TABLE banks (
  id      SERIAL,
  name    VARCHAR(50) UNIQUE,
  dollars INTEGER
);

INSERT INTO banks (name, dollars) VALUES ('test bank', 200);
`

var (
	testDB              *PostgresDB
	trader              Player
	receiver            Player
	traderTransaction   Transaction
	receiverTransaction Transaction
	offer               Offer
	trade               Trade
	err                 error
)

func NewTestDB() *PostgresDB {
	// TODO not sure if ssl should be enabled in tests
	dbConfigs := "user=rce_admin dbname=rce_test sslmode=disable"
	db, err := sql.Open("postgres", dbConfigs)

	if err != nil {
		log.Fatal("\n\n\nfailure connecting to database with configs:\n  ", dbConfigs, "\nerr:\n  ", err, "\n\n\n")
	}

	testDB := &PostgresDB{DB: db}
	testDB.Seed(testSeed)

	return testDB
}
