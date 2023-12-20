package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello world"))
	})

	apiRouter := initApiRouter()

	router.Mount("/api", apiRouter)

	return router
}

func initApiRouter() chi.Router {
	apiRouter := chi.NewRouter()

	transactionsRouter := initTransactionsRouter()
	apiRouter.Mount("/transactions", transactionsRouter)

	return apiRouter
}
