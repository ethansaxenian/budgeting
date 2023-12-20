package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/db"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
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
		http.Error(w, fmt.Sprintf("Transaction with id %d not found", id), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tr types.TransactionCreate
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	if tr.Date == "" {
		tr.Date = util.GetCurrentDate()
	}

	id, err := db.CreateTransaction(tr)
	if err != nil {
		http.Error(w, "Error creating transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var tr types.TransactionCreate
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	rowCount, err := db.UpdateTransaction(id, tr)
	if err != nil {
		http.Error(w, "Error updating transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Transaction with id %d not found", id), http.StatusBadRequest)
		return
	}

	rowCount, err := db.DeleteTransaction(id)
	if err != nil {
		http.Error(w, "Error deleting transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))

}