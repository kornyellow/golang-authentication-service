package main

import (
	"log"
	"os"

	"backend/internal/domain"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Env
	_ = godotenv.Load("../.env")

	// 2. Connect DB
	database.InitDB()

	// 3. Create Tables
	database.DB.AutoMigrate(&domain.User{})

	// 4. Dependency Injection
	userRepo := repository.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// 5. Setup Router
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", authHandler.GetProfile)
		}
	}

	// 6. Run
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🔥 Server running on port " + port)
	r.Run(":" + port)
}
