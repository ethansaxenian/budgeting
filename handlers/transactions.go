package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/db"
	"github.com/go-chi/chi/v5"
)

func GetTransactions(w http.ResponseWriter, _ *http.Request) {
	transactions, err := db.GetTransactions()
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction, err := db.GetTransactionByID(id)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}
