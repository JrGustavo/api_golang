package middleware

import (
	"github.com/JrGustavo/api_golang/auth"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareJSONAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := auth.ValidarToken(r)
		if err != nil {
			http.Error(w, "Falha na autenticação", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
