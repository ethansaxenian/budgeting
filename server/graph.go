package server

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/ethansaxenian/budgeting/components/graph"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/go-chi/chi/v5"
)

func getGraphData(transactions []database.Transaction, year int, month time.Month) types.GraphData {
	dayTotals := map[int]float64{}
	for _, t := range transactions {
		if t.Date.Month() == month && t.TransactionType == types.EXPENSE {
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
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid month ID", http.StatusBadRequest)
		return
	}

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	db := database.New(conn)

	month, err := db.GetMonthByID(ctx, monthID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthTransactions, err := db.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{ID: monthID, TransactionType: types.EXPENSE},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthDate, err := month.Date()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datasets := []types.GraphData{getGraphData(monthTransactions, month.Year, month.Month)}

	y, m, _ := monthDate.AddDate(0, -1, 0).Date()
	lastMonth, err := db.GetMonthByMonthAndYear(ctx, database.GetMonthByMonthAndYearParams{Month: m, Year: y})
	if err != nil {
		log.Printf("Failed to get last month (%s %d): %s", m, y, err)
	} else {
		lastMonthTransactions, err := db.GetTransactionsByMonthIDAndType(
			ctx,
			database.GetTransactionsByMonthIDAndTypeParams{ID: lastMonth.ID, TransactionType: types.EXPENSE},
		)
		if err != nil {
			log.Printf("Failed to get transactions for last month (%s %d): %s", m, y, err)
		} else {
			datasets = append(datasets, getGraphData(lastMonthTransactions, lastMonth.Year, lastMonth.Month))
		}
	}

	w.WriteHeader(http.StatusOK)
	graph.Graph(datasets).Render(ctx, w)
}
