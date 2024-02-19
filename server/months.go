package server

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/months"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleMonthShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	month, err := s.db.GetMonthByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allMonths, err := s.db.GetMonths()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(allMonths, func(i, j int) bool {
		return allMonths[i].Year > allMonths[j].Year || (allMonths[i].Year == allMonths[j].Year && allMonths[i].Month > allMonths[j].Month)
	})

	w.WriteHeader(http.StatusOK)
	months.MonthPage(month, allMonths).Render(r.Context(), w)
}
