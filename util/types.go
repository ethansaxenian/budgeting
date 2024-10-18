package util

import "github.com/ethansaxenian/budgeting/database"

var EXPENSE_CATEGORIES = []database.Category{
	database.CategoryFood,
	database.CategoryGifts,
	database.CategoryMedical,
	database.CategoryHome,
	database.CategoryTransportation,
	database.CategoryPersonal,
	database.CategorySavings,
	database.CategoryUtilities,
	database.CategoryTravel,
	database.CategoryOther,
}

var INCOME_CATEGORIES = []database.Category{
	database.CategoryPaycheck,
	database.CategoryBonus,
	database.CategoryInterest,
	database.CategoryGifts,
	database.CategoryOther,
	database.CategoryCashback,
}

var CATEGORIES_BY_TYPE = map[database.TransactionType][]database.Category{
	database.TransactionTypeExpense: EXPENSE_CATEGORIES,
	database.TransactionTypeIncome:  INCOME_CATEGORIES,
}

type GraphData struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}
