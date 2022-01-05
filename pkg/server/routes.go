package server

import (
	"net/http"

	"github.com/yonson2/go-rest-api-template/pkg/handlers"
)

func (s *Server) routes() {
	s.router.HandlerFunc("GET", "/", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {
	return handlers.Index()
}
