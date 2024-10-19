package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ethansaxenian/budgeting/assets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

type APIFunc func(conn *sql.Conn, w http.ResponseWriter, r *http.Request) error

func (s *Server) Handle(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		conn, err := s.db.Conn(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		if err := h(conn, w, r); err != nil {
			slog.Error("Error", "err", err.Error(), "method", r.Method, "path", r.URL, "remoteAddr", r.RemoteAddr)
			if apiErr, ok := err.(APIError); ok {
				http.Error(w, err.Error(), apiErr.StatusCode)
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
	}
}

func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Use(middleware.RedirectSlashes)

	assets.Mount(r)

	r.Get("/", s.Handle(base))
	r.Mount("/transactions", s.initTransactionsRouter())
	r.Mount("/months", s.initMonthsRouter())
	r.Mount("/budgets", s.initBudgetsRouter())

	return r
}

func (s *Server) initMonthsRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id:^[0-9]+}", s.Handle(HandleMonthShow))
	r.Get("/{id:^[0-9]+}/transactions/{transactionType:(income|expense)}", s.Handle(HandleTransactionsShow))
	r.Get("/{id:^[0-9]+}/budgets/{transactionType:(income|expense)}", s.Handle(HandleBudgetsShow))
	r.Get("/{id:^[0-9]+}/graph", s.Handle(HandleGraphShow))

	return r
}

func (s *Server) initTransactionsRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", s.Handle(HandleTransactionAdd))
	r.Put("/{id:^[0-9]+}", s.Handle(HandleTransactionEdit))
	r.Delete("/{id:^[0-9]+}", s.Handle(HandleTransactionDelete))

	return r
}

func (s *Server) initBudgetsRouter() chi.Router {
	r := chi.NewRouter()
	r.Patch("/{id:^[0-9]+}", s.Handle(HandleBudgetEdit))

	return r
}
