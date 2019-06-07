package main

import (
	"log"
	"weeny/apiserver"
	"weeny/application"
	"weeny/cache"
	"weeny/hasher"
)

func main() {

	s := apiserver.NewServer(
		application.New(
			cache.NewInMemory(),
			&hasher.Md5{},
		),
	)

	if err := s.Start("localhost", 8000); err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}
