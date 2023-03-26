package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"context"

	"github.com/canalesb93/user-server/internal/database"
)

// Server represents the main server struct.
type Server struct {
	db     *database.Database
	httpServer *http.Server
}

// NewServer creates a new instance of the server.
func NewServer(db *database.Database) *Server {
	return &Server{
		db: db,
		httpServer: &http.Server{
			Addr:         ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      nil, // will be set later
		},
	}
}

// Start starts the server.
func (s *Server) Start() {
	
	// Create a new ServeMux.
	mux := http.NewServeMux()

	// Register each handler function with its corresponding route.
	mux.HandleFunc("/user", s.getUserHandler)
	mux.HandleFunc("/", s.getUsersHandler)
	mux.HandleFunc("/users", s.getUsersHandler)
	mux.HandleFunc("/users/new", s.createUserHandler)

	s.httpServer.Handler = mux

	// Start the server.
	fmt.Printf("Server is listening on %s\n", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Could not start server: ", err)
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    fmt.Printf("Server shutting down...")
    return s.httpServer.Shutdown(ctx)
}
