package server

import (
	"net/http"

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

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("<h1>Hello, world!</h1>"))
	})

	apiRouter := s.initApiRouter()

	r.Mount("/api", apiRouter)

	return r
}

func (s *Server) initApiRouter() chi.Router {
	apiRouter := chi.NewRouter()

	apiRouter.Mount("/transactions", s.initTransactionsRouter())
	apiRouter.Mount("/months", s.initMonthsRouter())

	return apiRouter
}

func (s *Server) initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.HandleGetTransactions)
	r.Get("/{id:^[0-9]+}", s.HandleGetTransactionByID)
	r.Post("/", s.HandleCreateTransaction)
	r.Put("/{id:^[0-9]+}", s.HandleUpdateTransaction)
	r.Delete("/{id:^[0-9]+}", s.HandleDeleteTransaction)

	return r
}

func (s *Server) initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.HandleGetMonths)
	r.Get("/{id:^[0-9]+}", s.HandleGetMonthByID)
	r.Post("/", s.HandleCreateMonth)
	r.Put("/{id:^[0-9]+}", s.HandleUpdateMonth)
	r.Get("/{id:^[0-9]+}/transactions", s.HandleGetTransactionsByMonthID)

	return r
}
