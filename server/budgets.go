package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/budgets"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func HandleBudgetsShow(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid month ID"))
	}

	transactionType := database.TransactionType(chi.URLParam(r, "transactionType"))

	q := database.New(conn)

	allBudgetItems, err := q.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{
			MonthID:         monthID,
			TransactionType: transactionType,
		},
	)
	if err != nil {
		return err
	}

	budgetItems := []database.BudgetItem{}

	availableCategories := util.CATEGORIES_BY_TYPE[transactionType]

	for _, b := range allBudgetItems {
		if util.Includes(availableCategories, b.Category) {
			budgetItems = append(budgetItems, b)
		}
	}

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems, monthID, transactionType).Render(ctx, w)

	return nil
}

func HandleBudgetEdit(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid budget ID"))
	}

	amt, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid amount"))
	}

	q := database.New(conn)

	budget, err := q.PatchBudget(ctx, database.PatchBudgetParams{Amount: amt, ID: id})
	if err == sql.ErrNoRows {
		return NewAPIError(http.StatusNotFound, fmt.Errorf("budget with ID %d not found", id))
	} else if err != nil {
		return err
	}

	allBudgetItems, err := q.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{
			MonthID:         budget.MonthID,
			TransactionType: budget.TransactionType,
		},
	)
	if err != nil {
		return err
	}

	budgetItems := []database.BudgetItem{}

	availableCategories := util.CATEGORIES_BY_TYPE[budget.TransactionType]

	for _, b := range allBudgetItems {
		if util.Includes(availableCategories, b.Category) {
			budgetItems = append(budgetItems, b)
		}
	}

	w.WriteHeader(http.StatusOK)
	budgets.BudgetTable(budgetItems, budget.MonthID, budget.TransactionType).Render(ctx, w)

	return nil
}
