package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/maragudk/snorkel"

	"github.com/maragudk/ihukom/sql"
)

type Server struct {
	db     *sql.Database
	log    *snorkel.Logger
	mux    *chi.Mux
	server *http.Server
}

type NewServerOptions struct {
	DB  *sql.Database
	Log *snorkel.Logger
}

func NewServer(opts NewServerOptions) *Server {
	mux := chi.NewMux()
	return &Server{
		db:  opts.DB,
		log: opts.Log,
		mux: mux,
		server: &http.Server{
			Addr:              ":8080",
			Handler:           mux,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	s.log.Event("Starting http server", 1)

	s.setupRoutes()

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.log.Event("Stopping http server", 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	s.log.Event("Stopped http server", 1)
	return nil
}
