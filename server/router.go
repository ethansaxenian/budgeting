package server

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"sort"
	"time"

	"github.com/ethansaxenian/budgeting/assets"
	"github.com/ethansaxenian/budgeting/components/layout"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Use(middleware.RedirectSlashes)

	assets.Mount(r)

	r.Get("/", s.baseHandler)
	r.Mount("/transactions", s.initTransactionsRouter())
	r.Mount("/months", s.initMonthsRouter())
	r.Mount("/budgets", s.initBudgetsRouter())

	return r
}

func createCurrMonth(db *database.DB) (types.Month, error) {
	m := time.Now().Month()
	y := time.Now().Year()

	id, err := db.CreateMonth(types.MonthCreate{Month: m, Year: y})
	if err != nil {
		return types.Month{}, err
	}

	if err = db.CreateNewBudgetsForMonth(id); err != nil {
		return types.Month{}, err
	}

	currMonth := types.Month{ID: id, Month: m, Year: y}

	return currMonth, nil
}

func (s *Server) baseHandler(w http.ResponseWriter, r *http.Request) {
	currMonth, err := s.db.GetMonthByMonthAndYear(util.GetCurrMonthCtx(r.Context()))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		currMonth, err = createCurrMonth(s.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	allMonths, err := s.db.GetMonths()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(allMonths, func(i, j int) bool {
		return allMonths[i].Year > allMonths[j].Year || (allMonths[i].Year == allMonths[j].Year && allMonths[i].Month > allMonths[j].Month)
	})

	ctx := util.WithCurrMonthCtx(context.Background(), currMonth.FormatStr())
	layout.Base(allMonths).Render(ctx, w)
}

func (s *Server) initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id:^[0-9]+}", s.HandleMonthShow)

	return r
}

func (s *Server) initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.HandleTransactionsShow)
	r.Post("/", s.HandleTransactionAdd)
	r.Put("/{id:^[0-9]+}", s.HandleTransactionEdit)
	r.Delete("/{id:^[0-9]+}", s.HandleTransactionDelete)

	return r
}

func (s *Server) initBudgetsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.HandleBudgetsShow)
	r.Patch("/{id:^[0-9]+}", s.HandleBudgetEdit)

	return r
}
