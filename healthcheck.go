package main

import (
	"net/http"
)

func (s *HTTPServer) handleHealthcheck(w http.ResponseWriter, r *http.Request) error {
	resp := "OK"

	return WriteJSON(w, http.StatusOK, resp)
}
