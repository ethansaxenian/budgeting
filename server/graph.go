package server

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/ethansaxenian/budgeting/components/graph"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
)

func getGraphData(transactions []types.Transaction, year int, month time.Month) types.GraphData {
	dayTotals := map[int]float64{}
	for _, t := range transactions {
		if t.Date.Month() == month && t.Type == types.EXPENSE {
			dayTotals[t.Date.Day()] += t.Amount
		}
	}

	y, m, _ := time.Now().Date()
	var lastDay int
	if year == y && month == m {
		lastDay = time.Now().Day()
	} else {
		lastDay = time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()
	}

	amounts := []float64{dayTotals[1]}
	for day := 2; day <= lastDay; day++ {
		amounts = append(amounts, dayTotals[day]+dayTotals[day-1])
	}

	sort.Float64s(amounts)

	return types.GraphData{
		Label: fmt.Sprintf("%s %d", month.String(), year),
		Data:  amounts,
	}
}

func (s *Server) HandleGraphShow(w http.ResponseWriter, r *http.Request) {
	transactions, err := s.db.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	months, err := s.db.GetMonths()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datasets := []types.GraphData{}
	for _, m := range months {
		if (m.Month == util.CurrentMonth() || m.Month == util.CurrentMonth()-1) && m.Year == util.CurrentYear() {
			datasets = append(datasets, getGraphData(transactions, m.Year, m.Month))
		}
	}

	w.WriteHeader(http.StatusOK)
	graph.Graph(datasets).Render(r.Context(), w)
}
