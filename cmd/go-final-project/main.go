package main

import (
	"go-final-project/internal/db"
	"go-final-project/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbPath := os.Getenv("TODO_DBFILE")
	if dbPath == "" {
		dbPath = "scheduler.db"
	}

	err = db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}
