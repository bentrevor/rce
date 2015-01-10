package rce

type Currency string

const (
	Dollars Currency = "dollars"
)

const (
	InstitutionColumns string = "name,dollars"
)

type Institution interface {
	GetName() string
	GetDollars() int
	TableName() string
	ColumnNames() string
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

func (HedgeFund) ColumnNames() string {
	return "name,dollars"
}

func (Bank) ColumnNames() string {
	return "name,dollars"
}

// TODO the guide says you shouldn't have to make Get*() methods, but I don't
// know a better way to write methods with interface args
func (h HedgeFund) GetName() string {
	return h.Name
}

func (h HedgeFund) GetDollars() int {
	return h.Dollars
}

func (b Bank) GetName() string {
	return b.Name
}

func (b Bank) GetDollars() int {
	return b.Dollars
}
