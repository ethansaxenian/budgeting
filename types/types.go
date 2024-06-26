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

var CATEGORIES_BY_TYPE = map[TransactionType][]Category{
	EXPENSE: EXPENSE_CATEGORIES,
	INCOME:  INCOME_CATEGORIES,
}

type TransactionCreate struct {
	Date        time.Time
	Description string
	Category    Category
	Type        TransactionType
	Amount      float64
}

type TransactionUpdate TransactionCreate

type Transaction struct {
	Date        time.Time
	Description string
	Category    Category
	Type        TransactionType
	Amount      float64
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

func (m Month) HasDate(date time.Time) bool {
	return m.Month == date.Month() && m.Year == date.Year()
}

func (m Month) StartEndDates() (time.Time, time.Time) {
	start := time.Date(m.Year, m.Month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return start, end
}

func (m Month) Date() (time.Time, error) {
	return time.Parse("2006-01", m.FormatStr())
}

type BudgetCreate struct {
	Category Category
	Type     TransactionType
	MonthID  int
	Amount   float64
}

type BudgetUpdate BudgetCreate

type Budget struct {
	Category Category
	Type     TransactionType
	MonthID  int
	Amount   float64
	ID       int
}

type BudgetItem struct {
	Category Category
	Type     TransactionType
	ID       int
	Planned  float64
	Actual   float64
}

type GraphData struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}
