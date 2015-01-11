package rce

type Transaction struct {
	Trader   Player
	Receiver Player
	Amount   int
	Currency Currency
}

type Trade struct {
	Transactions []Transaction
	Desc         string
}

type Transactor struct{}

func (transactor Transactor) Execute(transaction Transaction, db *PostgresDB) error {
	traderUpdates := Update{}.Statements(transaction)
	receiverUpdates := Update{}.Statements(transaction)

	for _, update := range append(traderUpdates, receiverUpdates...) {
		db.Exec(update)
	}

	return nil
}
