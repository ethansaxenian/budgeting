package router

import (
	"net/http"

	"github.com/ethansaxenian/budgeting/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello world"))
	})

	apiRouter := initApiRouter()

	r.Mount("/api", apiRouter)

	return r
}

func initApiRouter() chi.Router {
	apiRouter := chi.NewRouter()

	transactionsRouter := initTransactionsRouter()
	apiRouter.Mount("/transactions", transactionsRouter)

	return apiRouter
}

func initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.GetTransactions)
	r.Get("/{id:^[0-9]+}", handlers.GetTransactionByID)
	r.Post("/", handlers.CreateTransaction)
	r.Put("/{id:^[0-9]+}", handlers.UpdateTransaction)
	r.Delete("/{id:^[0-9]+}", handlers.DeleteTransaction)

	return r
}