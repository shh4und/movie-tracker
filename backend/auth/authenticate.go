package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if exp, ok := claims["expiresAt"].(float64); ok {
			fmt.Printf("Expiration Time: %v, Current Time: %v\n", time.Unix(int64(exp), 0), time.Now())

			if time.Unix(int64(exp), 0).Before(time.Now()) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				ctx.Abort()
				return
			}
		}

		ctx.Set("userID", claims["userID"])
		ctx.Set("validUser", true)

		ctx.Next()
	}
}
