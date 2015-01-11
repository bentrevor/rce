package rce

type Currency string

const (
	Dollars Currency = "dollars"
)

const (
	PlayerColumnNames string = "name,dollars"
)

type Player struct {
	Name        string
	Dollars     int
	TableName   string
	ColumnNames string
}

func NewBank(name string) Player {
	return Player{
		Name:        name,
		Dollars:     200,
		TableName:   "banks",
		ColumnNames: PlayerColumnNames,
	}
}

func NewHedgeFund(name string) Player {
	return Player{
		Name:        name,
		Dollars:     100,
		TableName:   "hedge_funds",
		ColumnNames: PlayerColumnNames,
	}
}
