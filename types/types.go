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
