package fmtapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yokoe/gofmt-api-server/internal/pkgs/formatter"
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
		if len(result) == 0 {
			fmt.Fprintln(w, err)
		} else {
			fmt.Fprintln(w, result)
		}
		return
	}
	fmt.Fprintf(w, result)
}

func processFormatRequest(r *http.Request) (string, error) {
	if r.Method != "POST" {
		return "", fmt.Errorf("unsupported method %s", r.Method)
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	body := string(buf)

	if len(body) == 0 {
		return "", fmt.Errorf("no body")
	}

	return formatter.Format(body)
}
