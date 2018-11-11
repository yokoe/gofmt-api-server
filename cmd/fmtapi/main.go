package main

import (
	"log"

	"github.com/yokoe/gofmt-api-server/internal/app/fmtapi"
)

func main() {
	s := fmtapi.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
