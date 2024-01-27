package server

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/budgets"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
)

func (s *Server) HandleBudgetsShow(w http.ResponseWriter, r *http.Request) {
	monthID, err := strconv.Atoi(r.URL.Query().Get("month_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthBudgets, err := s.db.GetBudgets(monthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allTransactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	month, err := s.db.GetMonthByID(monthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transactionType := r.URL.Query().Get("type")

	expensesByCategory := map[types.Category]float64{}
	for _, t := range allTransactions {
		if util.GetDateMonth(t.Date) == month.FormatStr() && t.Type == types.EXPENSE {
			expensesByCategory[t.Category] += t.Amount
		}
	}

	incomeByCategory := map[types.Category]float64{}
	for _, t := range allTransactions {
		if util.GetDateMonth(t.Date) == month.FormatStr() && t.Type == types.INCOME {
			incomeByCategory[t.Category] += t.Amount
		}
	}

	typeSums := map[types.TransactionType]map[types.Category]float64{}
	typeSums[types.EXPENSE] = expensesByCategory
	typeSums[types.INCOME] = incomeByCategory

	budgetItems := []types.BudgetItem{}
	for _, b := range monthBudgets {
		if b.Type != types.TransactionType(transactionType) {
			continue
		}

		if b.Type == types.INCOME && !util.Includes[types.Category](types.INCOME_CATEGORIES, b.Category) {
			continue
		}

		if b.Type == types.EXPENSE && !util.Includes[types.Category](types.EXPENSE_CATEGORIES, b.Category) {
			continue
		}

		budgetItems = append(budgetItems, types.BudgetItem{
			ID:       b.ID,
			Category: b.Category,
			Planned:  b.Amount,
			Actual:   typeSums[b.Type][b.Category],
		})
	}

	sort.Slice(budgetItems, func(i, j int) bool {
		return budgetItems[i].Category < budgetItems[j].Category
	})

	ctx := util.WithTransactionTypeCtx(r.Context(), transactionType)
	ctx = util.WithCurrMonthIDCtx(ctx, monthID)

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems).Render(ctx, w)
}
