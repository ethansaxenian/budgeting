package routers

import (
	"github.com/ethansaxenian/budgeting/handlers"
	"github.com/go-chi/chi/v5"
)

func initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.GetTransactions)
	r.Get("/{id:^[0-9]+}", handlers.GetTransactionByID)

	return r
}
