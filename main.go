package main

import (
	"log"
	"tokogo/config"
	"tokogo/handlers"
	"tokogo/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	config.InitDB()

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()

	// Public routes (tidak perlu authentication)
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	// Protected routes (perlu authentication)
	protected := r.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Auth protected routes
		auth := protected.Group("/auth")
		{
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/profile", authHandler.GetProfile)
		}

		// Admin routes (perlu admin role)
		admin := protected.Group("/admin")
		admin.Use(middlewares.AdminMiddleware())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Welcome to admin dashboard",
					"user_id": c.GetUint("user_id"),
				})
			})
		}
	}

	// Start server
	port := config.GetEnv("SERVER_PORT", "8080")
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
