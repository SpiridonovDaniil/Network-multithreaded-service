package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"diploma/internal/domain"
)

type service interface {
	GetResultData() domain.ResultSetT
}

type Server struct {
	service      service
	server       http.Server
	shutDownTime time.Duration
}

func NewServer(service service, shutDownTime time.Duration) *Server {
	r := mux.NewRouter()
	r.HandleFunc("/api", HandleConnection(service))

	s := Server{
		service:      service,
		server:       http.Server{Addr: ":8282", Handler: r},
		shutDownTime: shutDownTime,
	}

	return &s
}

func (s *Server) Serve() error {
	return s.server.ListenAndServe()
}

func (s *Server) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutDownTime)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("gracefull shutdown не успешен, error: %w", err)
	}

	return nil
}
