package handlers

import (
	"encoding/json"
	"net/http"
)

type customError struct {
	Err string `json:"error"`
}

func JsonNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(customError{Err: "page not found"})
}
