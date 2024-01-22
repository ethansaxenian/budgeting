package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/months"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleMonthShow(w http.ResponseWriter, r *http.Request) {
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

	allTransactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}

	monthTransactions := []types.Transaction{}
	for _, tr := range allTransactions {
		if tr.Date.Month() == month.Month && tr.Date.Year() == month.Year {
			monthTransactions = append(monthTransactions, tr)
		}
	}

	ctx := util.WithCurrMonthCtx(r.Context(), month.FormatStr())
	w.WriteHeader(http.StatusOK)
	months.MonthPage(month, monthTransactions).Render(ctx, w)
}
