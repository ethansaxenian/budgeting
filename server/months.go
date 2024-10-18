package server

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/months"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleMonthShow(w http.ResponseWriter, r *http.Request) {
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

	db := database.New(conn)

	month, err := db.GetMonthByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allMonths, err := db.GetAllMonths(ctx)
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
