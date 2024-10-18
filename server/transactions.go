package server

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ethansaxenian/budgeting/components/transactions"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/types"
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

func (s *Server) HandleTransactionsShow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionType := types.TransactionType(chi.URLParam(r, "transactionType"))

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	db := database.New(conn)

	month, err := db.GetMonthByID(ctx, monthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startDate, endDate := month.StartEndDates()
	monthTransactions, err := db.GetTransactionsByTypeInDateRange(
		ctx,
		database.GetTransactionsByTypeInDateRangeParams{
			TransactionType: transactionType,
			StartDate:       startDate,
			EndDate:         endDate,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
}

func (s *Server) HandleTransactionEdit(w http.ResponseWriter, r *http.Request) {
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

	date, err := util.ParseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	desc := r.FormValue("description")
	cat := types.Category(r.FormValue("category"))

	newTransaction := types.TransactionUpdate{
		Description: desc,
		Amount:      amt,
		Date:        date,
		Category:    cat,
	}

	if err = s.db.UpdateTransaction(id, newTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := types.Transaction{
		Description: desc,
		Amount:      amt,
		Date:        date,
		Category:    cat,
		ID:          id,
	}

	w.Header().Set("HX-Trigger", "editTransaction")
	w.WriteHeader(http.StatusOK)
	transactions.TransactionRow(t).Render(context.Background(), w)
}

func (s *Server) HandleTransactionDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = s.db.DeleteTransaction(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "deleteTransaction")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandleTransactionAdd(w http.ResponseWriter, r *http.Request) {
	amt, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := util.ParseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTransaction := types.TransactionCreate{
		Description: r.FormValue("description"),
		Amount:      amt,
		Date:        date,
		Category:    types.Category(r.FormValue("category")),
		Type:        types.TransactionType(r.FormValue("type")),
	}

	if err := s.db.CreateTransaction(newTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "newTransaction")
	w.WriteHeader(http.StatusNoContent)
}
