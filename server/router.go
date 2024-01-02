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
	r.Get("/", s.GetTransactionsHandler)
	r.Get("/{id:^[0-9]+}", s.GetTransactionByIDHandler)
	r.Post("/", s.CreateTransactionHandler)
	r.Put("/{id:^[0-9]+}", s.UpdateTransactionHandler)
	r.Delete("/{id:^[0-9]+}", s.DeleteTransactionHandler)

	return r
}

func (s *Server) initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.GetMonthsHandler)
	r.Get("/{id:^[0-9]+}", s.GetMonthByIDHandler)
	r.Post("/", s.CreateMonthHandler)
	r.Put("/{id:^[0-9]+}", s.UpdateMonthHandler)
	r.Get("/{id:^[0-9]+}/transactions", s.GetTransactionsByMonthIDHandler)

	return r
}
