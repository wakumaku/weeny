package main

import (
	"log"
	"weeny/apiserver"
	"weeny/application"
	"weeny/cache"
	"weeny/hasher"
)

func main() {

	// c, err := cache.NewRedis("localhost", 6379)
	c, err := cache.NewInMemory()
	if err != nil {
		log.Fatalf("error while setting cache: %s", err)
	}

	app := application.New(c, &hasher.Md5{})

	server := apiserver.NewServer(app)

	if err := server.Start("localhost", 8000); err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}
