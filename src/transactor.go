package rce

import "fmt"

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

func (transactor Transactor) Execute(trade Trade) error {
	fmt.Println("executing: %s", trade.Desc)
	return nil
}
