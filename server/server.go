package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/ethansaxenian/budgeting/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Server struct {
	db     *sql.DB
	port   int
	server *http.Server
}

func (s *Server) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	if err := s.server.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func NewServer(port int, databaseURL string) (*Server, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
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

	s.server = server

	return s, nil
}
