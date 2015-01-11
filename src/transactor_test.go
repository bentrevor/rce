package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func NewTestTrade() Trade {
	trader = NewHedgeFund("test hedge fund")
	receiver = NewBank("test bank")

	traderTransaction = Transaction{Amount: 10, Currency: Dollars}
	receiverTransaction = Transaction{Amount: 5, Currency: Dollars}

	offer = Offer{TraderTransaction: traderTransaction, ReceiverTransaction: receiverTransaction}
	return Trade{
		Trader:   trader,
		Receiver: receiver,
		Offer:    offer,
		Desc:     "test trade",
	}
}

func TestTransaction_HasADescription(t *testing.T) {
	trade := NewTestTrade()
	assertEquals(t, trade.Desc, "test trade")
}

// func TestTransactor_CanExecuteATrade(t *testing.T) {
// 	transactor := Transactor{}
// 	trade := NewTestTrade()
// 	db := NewTestDB()
// 	transactor.Execute(trade, db)

// 	fmt.Println(trade.Offer.TraderTransaction)
// 	fmt.Println(trade.Offer.ReceiverTransaction)
// 	assertEquals(t, 95, db.GetBalance(trader)[Dollars])
// 	assertEquals(t, 205, db.GetBalance(receiver)[Dollars])
// }
