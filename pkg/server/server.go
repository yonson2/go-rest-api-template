package server

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yonson2/go-rest-api-template/pkg/config"
	"github.com/yonson2/go-rest-api-template/pkg/handlers"
	"github.com/yonson2/go-rest-api-template/pkg/storage"
)

type Server struct {
	router *httprouter.Router
	db     storage.Storage
	srv    http.Server
}

// New returns a new server instance.
func New(db storage.Storage) *Server {
	config := config.New()

	s := &Server{
		router: httprouter.New(),
	}
	s.db = db
	s.routes()
	s.router.NotFound = http.HandlerFunc(handlers.JSONNotFound)

	s.srv = http.Server{
		Addr:    ":" + config.Port,
		Handler: s.router,
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// StartServer starts listening on a TCP network, after it returns
// the server is ready to start receiving connections.
func (s *Server) StartServer() error {
	return s.srv.ListenAndServe()
}

// Shutdown function, mainly used when implementing graceful shutdown.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
func (s *Server) Close() error {
	return s.srv.Close()
}
