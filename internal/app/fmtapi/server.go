package fmtapi

import (
	"bytes"
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
	http.HandleFunc("/format", formatHandler)

	return http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fmtapi server")
}

func formatHandler(w http.ResponseWriter, r *http.Request) {
	result, err := processFormatRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintf(w, result)
}

func processFormatRequest(r *http.Request) (string, error) {
	if r.Method != "POST" {
		return "", fmt.Errorf("unsupported method %s", r.Method)
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return "", err
	}
	body := buf.String()

	if len(body) == 0 {
		return "", fmt.Errorf("no body")
	}

	return body, nil
}
