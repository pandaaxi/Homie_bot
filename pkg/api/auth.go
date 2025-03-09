package auth

import (
	"net/http"
	"os"
)

// AuthMiddleware checks the incoming request for the correct secret token.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Secret-Token")
		if token == "" || token != os.Getenv("AGENT_SECRET_TOKEN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
