package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

type Server struct {
	port int
	db   *database.DB
}

func NewServer() (*http.Server, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}

	s := &Server{
		port: port,
		db:   db,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: s.InitRouter(),
		BaseContext: func(_ net.Listener) context.Context {
			ctx := context.Background()
			ctx = util.WithNextSortCtx(ctx, util.GetNextSortCtx(ctx))
			return ctx
		},
	}

	return server, nil
}
