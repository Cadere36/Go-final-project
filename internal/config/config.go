package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBFile   string
	Password string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found, using environment variables")
	}
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}
	password := os.Getenv("TODO_PASSWORD")

	return Config{
		Port:     port,
		DBFile:   dbFile,
		Password: password,
	}
}
