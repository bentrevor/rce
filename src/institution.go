package rcfe

type Currency string

const (
	Dollars Currency = "dollars"
)

type Institution interface {
	GetName() string
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

// TODO the guide says you shouldn't have to make Get*() methods, but I don't
// know a better way to write methods with interface args
func (h HedgeFund) GetName() string {
	return h.Name
}

func (b Bank) GetName() string {
	return b.Name
}
