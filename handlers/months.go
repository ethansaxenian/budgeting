package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/db"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/go-chi/chi/v5"
)

func GetMonths(w http.ResponseWriter, _ *http.Request) {
	months, err := db.GetMonths()
	if err != nil {
		http.Error(w, "Error retrieving months", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(months)
}

func GetMonthByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid month ID", http.StatusBadRequest)
		return
	}

	month, err := db.GetMonthByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Month with id %d not found", id), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(month)
}

func CreateMonth(w http.ResponseWriter, r *http.Request) {
	var month types.MonthCreate
	if err := json.NewDecoder(r.Body).Decode(&month); err != nil {
		http.Error(w, "Invalid month data", http.StatusBadRequest)
		return
	}

	rowCount, err := db.CreateMonth(month)
	if err != nil {
		http.Error(w, "Error creating month", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(rowCount)))
}

func UpdateMonth(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid month ID", http.StatusBadRequest)
		return
	}

	var month types.MonthUpdate
	if err := json.NewDecoder(r.Body).Decode(&month); err != nil {
		http.Error(w, "Invalid month data", http.StatusBadRequest)
		return
	}

	rowCount, err := db.UpdateMonth(id, month)
	if err != nil {
		http.Error(w, "Error updating month", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))
}
