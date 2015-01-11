package rce

import (
	"fmt"
	"log"
)

type Transaction struct {
	Amount   int
	Currency Currency
}

type Offer struct {
	TraderTransaction   Transaction
	ReceiverTransaction Transaction
}

type Trade struct {
	Trader   Player
	Receiver Player
	Offer    Offer
	Desc     string
}

type Transactor struct{}

func (transactor Transactor) Execute(trade Trade, db *PostgresDB) error {
	traderUpdates := UpdateStatements(trade)
	receiverUpdates := UpdateStatements(trade)

	for _, update := range append(traderUpdates, receiverUpdates...) {
		db.Exec(update)
		rows, err := db.Query("select name,dollars from hedge_funds;")
		defer rows.Close()

		if err != nil {
			log.Fatal("got an error: ", err)
		}

		for rows.Next() {
			var name string
			var dollars int
			rows.Scan(&name, &dollars)

			fmt.Println("name: ", name, "dollars: ", dollars)
		}
	}
	// db.Exec(traderUpdate)
	// db.Exec(receiverUpdate)

	return nil
}
