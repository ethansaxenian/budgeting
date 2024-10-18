package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/budgets"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleBudgetsShow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionType := database.TransactionType(chi.URLParam(r, "transactionType"))

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	db := database.New(conn)

	allBudgetItems, err := db.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{
			MonthID:         monthID,
			TransactionType: transactionType,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(allBudgetItems)

	budgetItems := []database.BudgetItem{}

	availableCategories := types.CATEGORIES_BY_TYPE[transactionType]
	fmt.Println(availableCategories)

	for _, b := range allBudgetItems {
		if util.Includes(availableCategories, b.Category) {
			budgetItems = append(budgetItems, b)
		}
	}
	fmt.Println(budgetItems)

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems, monthID, transactionType).Render(ctx, w)
}

func (s *Server) HandleBudgetEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	allBudgetItems, err := db.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{
			MonthID:         budget.MonthID,
			TransactionType: budget.TransactionType,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	budgetItems := []database.BudgetItem{}

	availableCategories := types.CATEGORIES_BY_TYPE[budget.TransactionType]

	for _, b := range allBudgetItems {
		if util.Includes(availableCategories, b.Category) {
			budgetItems = append(budgetItems, b)
		}
	}

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems, budget.MonthID, budget.TransactionType).Render(ctx, w)
}
