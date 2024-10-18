package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Server struct {
	db   *sql.DB
	port int
}

func noop() error { return nil }

func NewServer() (*http.Server, func() error, error) {
	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		return nil, noop, err
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}

	s := &Server{
		db:   db,
		port: port,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.InitRouter(),
		BaseContext: func(_ net.Listener) context.Context {
			ctx := context.Background()
			ctx = util.WithNextSortCtx(ctx, util.GetNextSortCtx(ctx))
			return ctx
		},
	}

	close := func() error {
		if err := db.Close(); err != nil {
			return err
		}

		if err := server.Close(); err != nil {
			return err
		}

		return nil
	}

	return server, close, nil
}
