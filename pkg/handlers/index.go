package handlers

import (
	"fmt"
	"net/http"
)

// Index Handler
func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}
}
