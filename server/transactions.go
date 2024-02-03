package server

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ethansaxenian/budgeting/components/transactions"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func sortTransactions(transactionSlice []types.Transaction, sortParam string) {
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
	allTransactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	month := r.URL.Query().Get("month")
	transactionType := r.URL.Query().Get("type")

	filteredTransactions := []types.Transaction{}
	for _, tr := range allTransactions {
		if (util.GetDateMonth(tr.Date) == month || month == "") && (string(tr.Type) == transactionType || transactionType == "") {
			filteredTransactions = append(filteredTransactions, tr)
		}
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "" {
		sortParam = "date" + util.GetNextSortCtx(r.Context())
	}

	sortTransactions(filteredTransactions, sortParam)

	var dir string
	if strings.HasSuffix(sortParam, util.ContextValueSortDirDesc) {
		dir = util.ContextValueSortDirDesc
	} else {
		dir = util.ContextValueSortDirAsc
	}
	ctx := util.WithNextSortCtx(r.Context(), dir)
	ctx = util.WithCurrMonthCtx(ctx, month)
	ctx = util.WithTransactionTypeCtx(ctx, transactionType)

	w.WriteHeader(http.StatusOK)
	transactions.TransactionTable(filteredTransactions).Render(ctx, w)
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
