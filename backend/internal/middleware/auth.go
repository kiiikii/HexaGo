package middleware

import (
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		var tokenString string

		if len(header) > 7 && header[:7] == "Bearer " {
			tokenString = header[7:]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing of Malformed Authorization header"})
			return
		}

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired Token",
			})
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
