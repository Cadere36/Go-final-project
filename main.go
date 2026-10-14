package main

import (
	"log"
	"os"

	"go-final-project/internal/config"
	"go-final-project/internal/db"
	"go-final-project/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.Load()

	err := db.Init(cfg.DBFile)
	if err != nil {
		return err
	}
	defer db.DB.Close()

	return server.RunServer(cfg)

}
