package rce

type Currency string

const (
	Dollars Currency = "dollars"
)

type Institution struct {
	Name        string
	Dollars     int
	TableName   string
	ColumnNames string
}

func NewBank(name string) Institution {
	return Institution{
		Name:        name,
		Dollars:     200,
		TableName:   "banks",
		ColumnNames: "name,dollars",
	}
}

func NewHedgeFund(name string) Institution {
	return Institution{
		Name:        name,
		Dollars:     100,
		TableName:   "hedge_funds",
		ColumnNames: "name,dollars",
	}
}
