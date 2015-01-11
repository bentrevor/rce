package rce

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
	}

	return nil
}
