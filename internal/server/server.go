package server

import (
	"fmt"
	"go-final-project/internal/api"
	"net/http"
	"os"
)

func RunServer() error {
	httpPort := os.Getenv("TODO_PORT")
	if httpPort == "" {
		httpPort = "7540"
	}

	http.Handle("/", http.FileServer(http.Dir("./web/")))

	api.Init()

	fmt.Println("Сервер запущен на http://localhost:" + httpPort)
	return http.ListenAndServe(":"+httpPort, nil)
}
