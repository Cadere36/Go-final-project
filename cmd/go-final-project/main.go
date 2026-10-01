package main

import (
	"go-final-project/internal/server"
	"log"
)

func main() {

	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
