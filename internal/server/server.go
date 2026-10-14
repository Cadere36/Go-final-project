package server

import (
	"fmt"
	"net/http"

	"go-final-project/internal/api"
	"go-final-project/internal/config"
)

func RunServer(cfg config.Config) error {

	http.Handle("/", http.FileServer(http.Dir("./web/")))

	api.Init(cfg.Password)

	fmt.Println("Сервер запущен на http://localhost:" + cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, nil)
}
