package middleware

import (
	"MovieAPI/response"
	"MovieAPI/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Token not found")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			response.Unauthorized(c, "Token is invalid or expired")
			c.Abort()
			return
		}

		c.Set("UserID", claims.UserID)
		c.Set("Email", claims.Email)
		c.Set("Role", claims.Role)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("Role")
		if !exists || role != "admin" {
			response.Forbidden(c, "Access denied, only admin roles are allowed")
			c.Abort()
			return
		}
		c.Next()
	}
}
