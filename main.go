package main

import (
	"log"
	"weeny/apiserver"
)

func main() {
	server := apiserver.NewServer()
	err := server.Start("localhost", 8000)
	if err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}
