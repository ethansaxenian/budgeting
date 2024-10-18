package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

func (s *Server) baseHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	currYear, currMonth, _ := util.CurrentDate()

	conn, err := s.db.Conn(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	db := database.New(tx)

	month, err := db.GetMonthByMonthAndYear(ctx, database.GetMonthByMonthAndYearParams{Month: currMonth, Year: currYear})

	switch err {
	case nil:
		break

	case sql.ErrNoRows:
		month, err = db.CreateMonth(ctx, database.CreateMonthParams{Month: currMonth, Year: currYear})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err = db.CreateNewBudgetsForMonth(ctx, month.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	http.Redirect(w, r, fmt.Sprintf("/months/%d", month.ID), http.StatusFound)
}
