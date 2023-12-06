package middleware

import (
	"booking_chi_text/pkg/auth"
	"context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}
		tokenManager, err := auth.NewManager("")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		user, err := tokenManager.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token or user not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
