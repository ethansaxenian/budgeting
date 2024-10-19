package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ethansaxenian/budgeting/components/transactions"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func sortTransactions(transactionSlice []database.Transaction, sortParam string) {
	switch sortParam {
	case "dateDesc":
		sort.Slice(transactionSlice, func(i, j int) bool {
			return transactionSlice[i].Date.After(transactionSlice[j].Date)
		})
	case "dateAsc":
		sort.Slice(transactionSlice, func(i, j int) bool {
			return transactionSlice[i].Date.Before(transactionSlice[j].Date)
		})
	case "amountDesc":
		sort.Slice(transactionSlice, func(i, j int) bool {
			return transactionSlice[i].Amount > transactionSlice[j].Amount
		})
	case "amountAsc":
		sort.Slice(transactionSlice, func(i, j int) bool {
			return transactionSlice[i].Amount < transactionSlice[j].Amount
		})
	}
}

func HandleTransactionsShow(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid month ID"))
	}

	transactionType := database.TransactionType(chi.URLParam(r, "transactionType"))

	q := database.New(conn)

	monthTransactions, err := q.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{
			ID:              monthID,
			TransactionType: transactionType,
		},
	)
	if err != nil {
		return err
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "" {
		sortParam = "date" + util.GetNextSortCtx(r.Context())
	}

	sortTransactions(monthTransactions, sortParam)

	var nextDir string
	if strings.HasSuffix(sortParam, util.ContextValueSortDirDesc) {
		nextDir = util.ContextValueSortDirDesc
	} else {
		nextDir = util.ContextValueSortDirAsc
	}

	ctx = util.WithNextSortCtx(ctx, nextDir)

	w.WriteHeader(http.StatusOK)
	transactions.TransactionTable(monthTransactions, monthID, transactionType).Render(ctx, w)

	return nil
}

func HandleTransactionEdit(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid transaction ID"))
	}

	amt, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid amount"))
	}

	date, err := util.ParseDate(r.FormValue("date"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid date"))
	}

	description := r.FormValue("description")
	category := database.Category(r.FormValue("category"))

	q := database.New(conn)

	t, err := q.UpdateTransaction(ctx, database.UpdateTransactionParams{
		Description: description,
		Amount:      amt,
		Date:        date,
		Category:    category,
		ID:          id,
	})
	if err == sql.ErrNoRows {
		return NewAPIError(http.StatusNotFound, fmt.Errorf("transaction with ID %d not found", id))
	} else if err != nil {
		return err
	}

	w.Header().Set("HX-Trigger", "editTransaction")
	w.WriteHeader(http.StatusOK)
	transactions.TransactionRow(t).Render(ctx, w)

	return nil
}

func HandleTransactionDelete(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid transaction ID"))
	}

	q := database.New(conn)

	if err = q.DeleteTransaction(ctx, id); err == sql.ErrNoRows {
		return NewAPIError(http.StatusNotFound, fmt.Errorf("transaction with ID %d not found", id))
	} else if err != nil {
		return err
	}

	w.Header().Set("HX-Trigger", "deleteTransaction")
	w.WriteHeader(http.StatusOK)

	return nil
}

func HandleTransactionAdd(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	amt, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid amount"))
	}

	date, err := util.ParseDate(r.FormValue("date"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid date"))
	}

	q := database.New(conn)

	newTransaction := database.CreateTransactionParams{
		Description:     r.FormValue("description"),
		Amount:          amt,
		Date:            date,
		Category:        database.Category(r.FormValue("category")),
		TransactionType: database.TransactionType(r.FormValue("type")),
	}

	if _, err := q.CreateTransaction(ctx, newTransaction); err != nil {
		return err
	}

	w.Header().Set("HX-Trigger", "newTransaction")
	w.WriteHeader(http.StatusNoContent)

	return nil
}
