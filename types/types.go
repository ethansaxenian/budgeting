package types

type TransactionType string

const (
	EXPENSE TransactionType = "expense"
	INCOME  TransactionType = "income"
)

type Category string

const (
	FOOD           Category = "food"
	GIFTS          Category = "gifts"
	MEDICAL        Category = "medical"
	HOME           Category = "home"
	TRANSPORTATION Category = "transportation"
	PERSONAL       Category = "personal"
	SAVINGS        Category = "savings"
	UTILITIES      Category = "utilities"
	TRAVEL         Category = "travel"
	OTHER          Category = "other"
	PAYCHECK       Category = "paycheck"
	BONUS          Category = "bonus"
	INTREST        Category = "intrest"
)

type Transaction struct {
	ID          int             `json:"id"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Date        string          `json:"date"`
	Category    Category        `json:"category"`
	Type        TransactionType `json:"type"`
	MonthID     int             `json:"month_id"`
}

type TransactionCreate struct {
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Date        string          `json:"date"`
	Category    Category        `json:"category"`
	Type        TransactionType `json:"type"`
	MonthID     int             `json:"month_id"`
}
