package main

import (
	"log"
	"weeny/apiserver"
	"weeny/application"
	"weeny/config"
)

func main() {

	s := apiserver.NewServer(
		application.New(
			config.CreateCacheFromConfig(),
			config.CreateHasherFromConfig(),
		),
	)

	if err := s.Start(8000); err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}
