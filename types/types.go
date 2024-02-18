package types

import (
	"fmt"
	"time"
)

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
	INTEREST       Category = "interest"
	CASHBACK       Category = "cashback"
)

var ALL_CATEGORIES = []Category{
	FOOD,
	GIFTS,
	MEDICAL,
	HOME,
	TRANSPORTATION,
	PERSONAL,
	SAVINGS,
	UTILITIES,
	TRAVEL,
	OTHER,
	PAYCHECK,
	BONUS,
	INTEREST,
	CASHBACK,
}

var EXPENSE_CATEGORIES = []Category{
	FOOD,
	GIFTS,
	MEDICAL,
	HOME,
	TRANSPORTATION,
	PERSONAL,
	SAVINGS,
	UTILITIES,
	TRAVEL,
	OTHER,
}

var INCOME_CATEGORIES = []Category{
	PAYCHECK,
	BONUS,
	INTEREST,
	GIFTS,
	OTHER,
	CASHBACK,
}

type TransactionCreate struct {
	Description string
	Amount      float64
	Date        time.Time
	Category    Category
	Type        TransactionType
}

type TransactionUpdate TransactionCreate

type Transaction struct {
	Description string
	Amount      float64
	Date        time.Time
	Category    Category
	Type        TransactionType
	ID          int
}

type MonthCreate struct {
	Month time.Month
	Year  int
}

type Month struct {
	ID    int
	Month time.Month
	Year  int
}

func (m Month) FormatStr() string {
	return fmt.Sprintf("%d-%02d", m.Year, m.Month)
}

type BudgetCreate struct {
	MonthID  int
	Category Category
	Amount   float64
	Type     TransactionType
}

type BudgetUpdate BudgetCreate

type Budget struct {
	MonthID  int
	Category Category
	Amount   float64
	Type     TransactionType
	ID       int
}

type BudgetItem struct {
	ID       int
	Category Category
	Planned  float64
	Actual   float64
}

type GraphData struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}
