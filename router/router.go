package router

import (
	"net/http"

	"github.com/ethansaxenian/budgeting/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Use(middleware.RedirectSlashes)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("<h1>Hello, world!</h1>"))
	})

	apiRouter := initApiRouter()

	r.Mount("/api", apiRouter)

	return r
}

func initApiRouter() chi.Router {
	apiRouter := chi.NewRouter()

	apiRouter.Mount("/transactions", initTransactionsRouter())
	apiRouter.Mount("/months", initMonthsRouter())

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

func initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.GetMonths)
	r.Get("/{id:^[0-9]+}", handlers.GetMonthByID)
	r.Post("/", handlers.CreateMonth)
	r.Put("/{id:^[0-9]+}", handlers.UpdateMonth)
	r.Get("/{id:^[0-9]+}/transactions", handlers.GetTransactionsByMonthID)

	return r
}
