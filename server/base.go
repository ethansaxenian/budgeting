package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

func base(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := database.New(tx)

	currYear, currMonth, _ := util.CurrentDate()

	month, err := q.GetMonthByMonthAndYear(ctx, database.GetMonthByMonthAndYearParams{Month: currMonth, Year: currYear})

	switch err {
	case nil:
		break

	case sql.ErrNoRows:
		month, err = q.CreateMonth(ctx, database.CreateMonthParams{Month: currMonth, Year: currYear})
		if err != nil {
			return err
		}

		if _, err = q.CreateNewBudgetsForMonth(ctx, month.ID); err != nil {
			return err
		}

	default:
		return err
	}

	tx.Commit()

	http.Redirect(w, r, fmt.Sprintf("/months/%d", month.ID), http.StatusFound)

	return nil
}
