package server

import (
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

func (s *Server) HandleTransactionsShow(w http.ResponseWriter, r *http.Request) {
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

	q := database.New(conn)

	monthTransactions, err := q.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{
			ID:              monthID,
			TransactionType: transactionType,
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

	date, err := util.ParseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	desc := r.FormValue("description")
	cat := database.Category(r.FormValue("category"))

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	q := database.New(conn)

	t, err := q.UpdateTransaction(ctx, database.UpdateTransactionParams{
		Description: desc,
		Amount:      amt,
		Date:        date,
		Category:    cat,
		ID:          id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "editTransaction")
	w.WriteHeader(http.StatusOK)
	transactions.TransactionRow(t).Render(ctx, w)
}

func (s *Server) HandleTransactionDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
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

	q := database.New(conn)

	if err = q.DeleteTransaction(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "deleteTransaction")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandleTransactionAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	q := database.New(conn)

	newTransaction := database.CreateTransactionParams{
		Description:     r.FormValue("description"),
		Amount:          amt,
		Date:            date,
		Category:        database.Category(r.FormValue("category")),
		TransactionType: database.TransactionType(r.FormValue("type")),
	}

	if _, err := q.CreateTransaction(ctx, newTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "newTransaction")
	w.WriteHeader(http.StatusNoContent)
}
