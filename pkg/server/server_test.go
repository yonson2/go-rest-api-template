package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/yonson2/go-rest-api-template/pkg/server"
)

func TestHandleIndex(t *testing.T) {
	is := is.New(t)
	srv := server.New(nil)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request %v", err)
	}
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	is.Equal(res.StatusCode, http.StatusOK)
}
