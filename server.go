package todo

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(httpServer *http.Server) error {
	s.httpServer = httpServer
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return s.httpServer.Shutdown(ctx)
}
