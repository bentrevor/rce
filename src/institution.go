package rcfe

type Institution interface {
	TableName() string
}

type HedgeFund struct {
	Name    string
	Dollars int
}

type Bank struct {
	Name    string
	Dollars int
}

func (HedgeFund) TableName() string {
	return "hedge_funds"
}

func (Bank) TableName() string {
	return "banks"
}
