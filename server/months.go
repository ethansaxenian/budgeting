package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/types"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleGetMonths(w http.ResponseWriter, _ *http.Request) {
	months, err := s.db.GetMonths()
	if err != nil {
		http.Error(w, "Error retrieving months", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(months)
}

func (s *Server) HandleGetMonthByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid month ID", http.StatusBadRequest)
		return
	}

	month, err := s.db.GetMonthByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Month with id %d not found", id), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(month)
}

func (s *Server) HandleCreateMonth(w http.ResponseWriter, r *http.Request) {
	var month types.MonthCreate
	if err := json.NewDecoder(r.Body).Decode(&month); err != nil {
		http.Error(w, "Invalid month data", http.StatusBadRequest)
		return
	}

	rowCount, err := s.db.CreateMonth(month)
	if err != nil {
		http.Error(w, "Error creating month", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(rowCount)))
}

func (s *Server) HandleUpdateMonth(w http.ResponseWriter, r *http.Request) {
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

	rowCount, err := s.db.UpdateMonth(id, month)
	if err != nil {
		http.Error(w, "Error updating month", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(rowCount)))
}
