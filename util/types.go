package util

import "github.com/ethansaxenian/budgeting/database"

var EXPENSE_CATEGORIES = []database.Category{
	database.CategoryFood,
	database.CategoryGifts,
	database.CategoryHome,
	database.CategoryMedical,
	database.CategoryOther,
	database.CategoryPersonal,
	database.CategorySavings,
	database.CategoryTransportation,
	database.CategoryTravel,
	database.CategoryUtilities,
}

var INCOME_CATEGORIES = []database.Category{
	database.CategoryBonus,
	database.CategoryCashback,
	database.CategoryGifts,
	database.CategoryInterest,
	database.CategoryOther,
	database.CategoryPaycheck,
}

var CATEGORIES_BY_TYPE = map[database.TransactionType][]database.Category{
	database.TransactionTypeExpense: EXPENSE_CATEGORIES,
	database.TransactionTypeIncome:  INCOME_CATEGORIES,
}

type GraphData struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}
