package fmtapi

import (
	"fmt"
	"net/http"
)

// Server represents api server
type Server struct {
}

// NewServer returns an instance of Server
func NewServer() *Server {
	return &Server{}
}

// Run starts listening
func (s *Server) Run() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fmtapi server")
}
