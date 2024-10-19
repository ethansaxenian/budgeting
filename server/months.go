package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethansaxenian/budgeting/components/months"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/go-chi/chi/v5"
)

func HandleMonthShow(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid month ID"))
	}

	q := database.New(conn)

	month, err := q.GetMonthByID(ctx, id)
	if err == sql.ErrNoRows {
		return NewAPIError(http.StatusNotFound, fmt.Errorf("month with ID %d not found", id))
	} else if err != nil {
		return err
	}

	allMonths, err := q.GetAllMonths(ctx)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	months.MonthPage(month, allMonths).Render(ctx, w)

	return nil
}
