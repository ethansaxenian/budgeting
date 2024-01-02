package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/database"
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

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 8080
	}

	s := &Server{
		port: port,
		db:   db,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.InitRouter(),
	}

	return server, nil
}
