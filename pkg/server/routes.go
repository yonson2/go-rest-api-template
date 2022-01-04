package server

import (
	"net/http"

	"github.com/yonson2/go-rest-api-template/pkg/handlers"
)

func (s *server) routes() {
	s.router.HandlerFunc("GET", "/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return handlers.Index()
}
