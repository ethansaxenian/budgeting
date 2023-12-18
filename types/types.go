package types

type Transaction struct {
	ID          int
	Description string
	Amount      float64
	Date        string
	Category    string
	Type        string
	MonthID     int
}
