package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/transactions"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleTransactionsShow(w http.ResponseWriter, r *http.Request) {
	allTransactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}
	transactions.TransactionTable(allTransactions).Render(context.Background(), w)
}

func (s *Server) HandleGetTransactions(w http.ResponseWriter, _ *http.Request) {
	transactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func (s *Server) HandleGetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction, err := s.db.GetTransactionByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Transaction with id %d not found", id), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func (s *Server) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tr types.TransactionCreate
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	if tr.Date == "" {
		tr.Date = util.GetCurrentDate()
	}

	id, err := s.db.CreateTransaction(tr)
	if err != nil {
		http.Error(w, "Error creating transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (s *Server) HandleUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var tr types.TransactionUpdate
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	rowCount, err := s.db.UpdateTransaction(id, tr)
	if err != nil {
		http.Error(w, "Error updating transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))
}

func (s *Server) HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Transaction with id %d not found", id), http.StatusBadRequest)
		return
	}

	rowCount, err := s.db.DeleteTransaction(id)
	if err != nil {
		http.Error(w, "Error deleting transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))
}

func (s *Server) HandleGetTransactionsByMonthID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Month with ID %d not found", id), http.StatusBadRequest)
		return
	}

	transactions, err := s.db.GetTransactionsByMonthID(id)
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}
