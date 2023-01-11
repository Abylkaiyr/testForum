package main

import (
	"log"

	"github.com/Abylkaiyr/forum/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	logger := apiserver.NewLogger()
	s := apiserver.New(config, logger)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	s.Server()
}
