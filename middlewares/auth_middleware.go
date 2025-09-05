package middlewares

import (
	"net/http"
	"strings"
	"tokogo/helpers"
	"tokogo/responses"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware untuk memvalidasi JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Error:   "unauthorized",
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Cek format Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Error:   "unauthorized",
				Message: "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validasi token
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Error:   "unauthorized",
				Message: "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user info ke context untuk digunakan di handler
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// AdminMiddleware middleware untuk memvalidasi admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cek apakah user sudah login (AuthMiddleware harus dipanggil dulu)
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Error:   "unauthorized",
				Message: "User not authenticated",
			})
			c.Abort()
			return
		}

		// Cek apakah user adalah admin
		if userRole != "admin" {
			c.JSON(http.StatusForbidden, responses.ErrorResponse{
				Error:   "forbidden",
				Message: "Admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
