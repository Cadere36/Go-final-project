package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeText(w http.ResponseWriter, status int, text string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Printf("write response: %v", err)
	}
}
