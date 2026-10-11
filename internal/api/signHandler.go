package api

import (
	"encoding/json"
	"go-final-project/internal/auth"
	"net/http"
	"os"
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var req struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "bad request"})
		return
	}

	pass := os.Getenv("TODO_PASSWORD")
	if pass == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "TODO_PASSWORD not set"})
		return
	}
	if req.Password != pass {
		writeJson(w, http.StatusUnauthorized, map[string]string{"error": "неверный пароль"})
		return
	}
	token, err := auth.GenerateToken(pass)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "token generation failed"})
		return
	}
	writeJson(w, http.StatusOK, map[string]string{"token": token})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var token string

			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}
			if !auth.ValidateToken(token, pass) {

				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
