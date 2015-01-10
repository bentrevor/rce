package rce_test

import (
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

var (
	trader              = NewHedgeFund("hf")
	receiver            = NewBank("b")
	traderTransaction   = Transaction{Amount: 10, Currency: Dollars}
	receiverTransaction = Transaction{Amount: 5, Currency: Dollars}
	offer               = Offer{TraderTransaction: traderTransaction, ReceiverTransaction: receiverTransaction}
	trade               = Trade{
		Trader:   trader,
		Receiver: receiver,
		Offer:    offer,
		Desc:     "test trade",
	}
)

func TestTransaction_HasADescription(t *testing.T) {
	assertEquals(t, trade.Desc, "test trade")
}
