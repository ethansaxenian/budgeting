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

	r.Mount("/transactions", s.initTransactionsRouter())

	// r.Mount("/api", s.initAPIRouter())

	return r
}

func (s *Server) initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.HandleTransactionsShow)
	r.Put("/{id:^[0-9]+}", s.HandleTransactionEdit)

	return r
}

// func (s *Server) initAPIRouter() chi.Router {
// 	apiRouter := chi.NewRouter()

// 	apiRouter.Mount("/transactions", s.initTransactionsAPIRouter())
// 	apiRouter.Mount("/months", s.initMonthsAPIRouter())

// 	return apiRouter
// }

// func (s *Server) initTransactionsAPIRouter() chi.Router {
// 	r := chi.NewRouter()
// 	r.Get("/", s.HandleGetTransactions)
// 	r.Get("/{id:^[0-9]+}", s.HandleGetTransactionByID)
// 	r.Post("/", s.HandleCreateTransaction)
// 	r.Put("/{id:^[0-9]+}", s.HandleUpdateTransaction)
// 	r.Delete("/{id:^[0-9]+}", s.HandleDeleteTransaction)

// 	return r
// }

// func (s *Server) initMonthsAPIRouter() chi.Router {
// 	r := chi.NewRouter()
// 	r.Get("/", s.HandleGetMonths)
// 	r.Get("/{id:^[0-9]+}", s.HandleGetMonthByID)
// 	r.Post("/", s.HandleCreateMonth)
// 	r.Put("/{id:^[0-9]+}", s.HandleUpdateMonth)
// 	r.Get("/{id:^[0-9]+}/transactions", s.HandleGetTransactionsByMonthID)

// 	return r
// }
