package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/ethansaxenian/budgeting/components/graph"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
)

func getGraphData(transactions []database.Transaction, year int, month time.Month) util.GraphData {
	dayTotals := map[int]float64{}
	for _, t := range transactions {
		if t.Date.Month() == month && t.TransactionType == database.TransactionTypeExpense {
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

	return util.GraphData{
		Label: fmt.Sprintf("%s %d", month.String(), year),
		Data:  amounts,
	}
}

func HandleGraphShow(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	monthID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid month ID"))
	}

	q := database.New(conn)

	month, err := q.GetMonthByID(ctx, monthID)
	if err == sql.ErrNoRows {
		return NewAPIError(http.StatusNotFound, fmt.Errorf("month with ID %d not found", monthID))
	} else if err != nil {
		return err
	}

	monthTransactions, err := q.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{ID: monthID, TransactionType: database.TransactionTypeExpense},
	)
	if err != nil {
		return err
	}

	monthDate, err := time.Parse("2006-01", fmt.Sprintf("%d-%02d", month.Year, month.Month))
	if err != nil {
		return err
	}

	datasets := []util.GraphData{getGraphData(monthTransactions, month.Year, month.Month)}

	y, m, _ := monthDate.AddDate(0, -1, 0).Date()
	lastMonth, err := q.GetMonthByMonthAndYear(ctx, database.GetMonthByMonthAndYearParams{Month: m, Year: y})
	if err != nil {
		slog.Error("Failed to get last month", "month", m, "year", y, "err", err)
	} else {
		lastMonthTransactions, err := q.GetTransactionsByMonthIDAndType(
			ctx,
			database.GetTransactionsByMonthIDAndTypeParams{ID: lastMonth.ID, TransactionType: database.TransactionTypeExpense},
		)
		if err != nil {
			slog.Error("Failed to get transactions for last month", "month", m, "year", y, "err", err)
		} else {
			datasets = append(datasets, getGraphData(lastMonthTransactions, lastMonth.Year, lastMonth.Month))
		}
	}

	w.WriteHeader(http.StatusOK)
	return graph.Graph(datasets).Render(ctx, w)
}
