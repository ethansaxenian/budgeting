package server

import (
	"github.com/ethansaxenian/budgeting/assets"
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

func (s *Server) initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id:^[0-9]+}", s.HandleMonthShow)
	r.Get("/{id:^[0-9]+}/transactions/{transactionType:(income|expense)}", s.HandleTransactionsShow)
	r.Get("/{id:^[0-9]+}/budgets/{transactionType:(income|expense)}", s.HandleBudgetsShow)
	r.Get("/{id:^[0-9]+}/graph", s.HandleGraphShow)

	return r
}

func (s *Server) initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", s.HandleTransactionAdd)
	r.Put("/{id:^[0-9]+}", s.HandleTransactionEdit)
	r.Delete("/{id:^[0-9]+}", s.HandleTransactionDelete)

	return r
}

func (s *Server) initBudgetsRouter() chi.Router {
	r := chi.NewRouter()
	r.Patch("/{id:^[0-9]+}", s.HandleBudgetEdit)

	return r
}
