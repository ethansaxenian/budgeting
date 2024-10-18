package server

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/budgets"
	"github.com/ethansaxenian/budgeting/database"
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

	ctx := r.Context()
	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	db := database.New(conn)

	monthBudgets, err := db.GetBudgetsByMonthIDAndType(ctx, database.GetBudgetsByMonthIDAndTypeParams{MonthID: monthID, TransactionType: transactionType})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	month, err := db.GetMonthByID(ctx, monthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startDate, endDate := month.StartEndDates()
	monthTransactions, err := db.GetTransactionsByTypeInDateRange(ctx, database.GetTransactionsByTypeInDateRangeParams{StartDate: startDate, EndDate: endDate})
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
		if !util.Includes(availableCategories, b.Category) {
			continue
		}

		budgetItems = append(budgetItems, types.BudgetItem{
			ID:       b.ID,
			Category: b.Category,
			Planned:  b.Amount,
			Actual:   byCategory[b.Category],
			Type:     b.TransactionType,
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

	ctx := r.Context()
	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	db := database.New(conn)

	budget, err := db.PatchBudget(ctx, database.PatchBudgetParams{Amount: amt, ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	month, err := db.GetMonthByID(ctx, budget.MonthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startDate, endDate := month.StartEndDates()

	monthTransactions, err := db.GetTransactionsByTypeAndCategoryInDateRange(
		ctx,
		database.GetTransactionsByTypeAndCategoryInDateRangeParams{
			TransactionType: budget.TransactionType,
			Category:        budget.Category,
			StartDate:       startDate,
			EndDate:         endDate,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var actual float64
	for _, t := range monthTransactions {
		actual += t.Amount
	}

	budgetItem := types.BudgetItem{
		ID:       budget.ID,
		Category: budget.Category,
		Planned:  budget.Amount,
		Actual:   actual,
		Type:     budget.TransactionType,
	}

	w.WriteHeader(http.StatusOK)
	budgets.BudgetRow(budgetItem).Render(r.Context(), w)
}
