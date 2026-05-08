package middleware

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
)

func AuthMiddleware(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet || r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			clientKey := r.Header.Get("X-API-Key")
			if clientKey != apiKey {
				handler.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized: Invalid API Key")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}


