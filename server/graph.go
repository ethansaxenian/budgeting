package server

import (
	"database/sql"
	"fmt"
	"net/http"
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
		if t.Date.Month() == month {
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

	total := 0.0
	amounts := []float64{dayTotals[0]}
	for day := 1; day <= lastDay; day++ {
		total += dayTotals[day]
		amounts = append(amounts, total)
	}

	return util.GraphData{
		Label: string(transactions[0].TransactionType),
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

	income, err := q.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{ID: monthID, TransactionType: database.TransactionTypeIncome},
	)
	if err != nil {
		return err
	}

	expenses, err := q.GetTransactionsByMonthIDAndType(
		ctx,
		database.GetTransactionsByMonthIDAndTypeParams{ID: month.ID, TransactionType: database.TransactionTypeExpense},
	)
	if err != nil {
		return err
	}

	datasets := []util.GraphData{
		getGraphData(income, month.Year, month.Month),
		getGraphData(expenses, month.Year, month.Month),
	}

	w.WriteHeader(http.StatusOK)
	return graph.Graph(fmt.Sprintf("%s %d", month.Month.String(), month.Year), datasets).Render(ctx, w)
}
