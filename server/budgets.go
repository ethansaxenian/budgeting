package server

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/budgets"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleBudgetsShow(w http.ResponseWriter, r *http.Request) {
	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionType := types.TransactionType(chi.URLParam(r, "transactionType"))

	monthBudgets, err := s.db.GetBudgetsByMonthID(monthID, transactionType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthTransactions, err := s.db.GetTransactionsByMonthID(monthID, transactionType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byCategory := map[types.Category]float64{}
	for _, t := range monthTransactions {
		byCategory[t.Category] += t.Amount
	}

	budgetItems := []types.BudgetItem{}

	availableCategories := types.CATEGORIES_BY_TYPE[transactionType]

	for _, b := range monthBudgets {
		if !util.Includes[types.Category](availableCategories, b.Category) {
			continue
		}

		budgetItems = append(budgetItems, types.BudgetItem{
			ID:       b.ID,
			Category: b.Category,
			Planned:  b.Amount,
			Actual:   byCategory[b.Category],
		})
	}

	sort.Slice(budgetItems, func(i, j int) bool {
		return budgetItems[i].Category < budgetItems[j].Category
	})

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems, monthID, transactionType).Render(r.Context(), w)
}

func (s *Server) HandleBudgetEdit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	amt, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = s.db.PatchBudget(id, amt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	budget, err := s.db.GetBudgetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allTransactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	month, err := s.db.GetMonthByID(budget.MonthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var actual float64
	for _, t := range allTransactions {
		if month.HasDate(t.Date) && t.Type == budget.Type && t.Category == budget.Category {
			actual += t.Amount
		}
	}

	budgetItem := types.BudgetItem{
		ID:       budget.ID,
		Category: budget.Category,
		Planned:  budget.Amount,
		Actual:   actual,
	}

	w.WriteHeader(http.StatusOK)
	ctx := util.WithTransactionTypeCtx(r.Context(), string(budget.Type))
	budgets.BudgetRow(budgetItem).Render(ctx, w)
}
