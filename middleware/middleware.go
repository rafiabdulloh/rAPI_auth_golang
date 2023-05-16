package middleware

import (
	"auth_with_token/auth"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// Memeriksa keberadaan token
	if tokenString == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Parsing token dan memeriksa validitas signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.PW), nil
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Memeriksa waktu kadaluarsa token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
		return
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		c.AbortWithStatusJSON(401, gin.H{"error": "Token has expired"})
		return
	}

}
