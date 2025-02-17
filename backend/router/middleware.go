package router

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shh4und/movie-tracker/auth"
)

// Authenticate is a middleware function for JWT authentication.
func Authenticate(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]
			claims := &auth.CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})

			if err != nil || !token.Valid {
				log.Printf("Token parsing error: %v", err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
				log.Printf("Token expired at: %v", claims.ExpiresAt.Time)
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			}

			log.Printf("Token valid for user ID: %d", claims.UserID)
			type contextKey string
			var userIDkey contextKey = "userID"
			ctx := context.WithValue(r.Context(), userIDkey, claims.Subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
