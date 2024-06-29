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
	db   *database.DB
	port int
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

	return server, nil
}
