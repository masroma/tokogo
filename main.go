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

	// Static file serving for uploaded images
	r.Static("/uploads", "./uploads")

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()
	categoryHandler := handlers.NewCategoryHandler()
	productHandler := handlers.NewProductHandler()
	userManagementHandler := handlers.NewUserManagementHandler()
	transactionHandler := handlers.NewTransactionHandler()

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

			categories := admin.Group("/categories")
			{
				categories.POST("", categoryHandler.CreateCategory)
				categories.GET("", categoryHandler.GetAllCategories)
				categories.GET("/:id", categoryHandler.GetCategoryByID)
				categories.PUT("/:id", categoryHandler.UpdateCategory)
				categories.DELETE("/:id", categoryHandler.DeleteCategory)
			}

			products := admin.Group("/products")
			{
				products.POST("", productHandler.CreateProduct)
				products.GET("", productHandler.GetAllProducts)
				products.GET("/:id", productHandler.GetProductByID)
				products.PUT("/:id", productHandler.UpdateProduct)
				products.DELETE("/:id", productHandler.DeleteProduct)
				products.GET("/categories/:category_id", productHandler.GetProductsByCategory)
			}

			userManagement := admin.Group("/user-management")
			{
				userManagement.POST("", userManagementHandler.CreateUser)
				userManagement.GET("", userManagementHandler.GetAllUsers)
				userManagement.GET("/:id", userManagementHandler.GetUserByID)
				userManagement.PUT("/:id", userManagementHandler.UpdateUser)
				userManagement.DELETE("/:id", userManagementHandler.DeleteUser)
			}

			transactions := admin.Group("/transactions")
			{
				transactions.GET("", transactionHandler.GetAllTransactions)
				transactions.GET("/:id", transactionHandler.GetTransactionByID)
				transactions.PUT("/:id/status", transactionHandler.UpdateTransactionStatus)
			}
		}
	}

	// Start server
	port := config.GetEnv("SERVER_PORT", "8080")
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
