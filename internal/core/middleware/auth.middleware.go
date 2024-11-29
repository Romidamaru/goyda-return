package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/core/config"
	"simple-api/internal/utils"
	"strings"
)

// AuthMiddleware validates the access token and extracts the userId
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Remove 'Bearer ' prefix
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { // Если префикс отсутствует
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Validate token and extract claims (userId)
		claims, err := utils.ValidateToken(token, config.Inst().SecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract userId from claims (convert from float64 to uint)
		userID, ok := claims["userId"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Store userId in context for further use
		c.Set("userID", uint(userID))
		c.Next()
	}
}
