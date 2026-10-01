package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func RunServer() error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	httpPort := os.Getenv("TODO_PORT")
	if httpPort == "" {
		httpPort = "7540"
	}

	http.Handle("/", http.FileServer(http.Dir("./web/")))

	fmt.Println("Сервер запущен на http://localhost:" + httpPort)
	return http.ListenAndServe(":"+httpPort, nil)

}
