package main

import (
	"log"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/routes"
	"backend/internal/services"

	_ "github.com/gin-contrib/cors"
)

func main() {
	cfg := config.Load()

	// Initialize database
	database.Init(cfg.DatabaseURL)

	// Defer closing the database connection
	defer database.GetDB().Close()

	authService := services.NewAuthService(cfg.JWTSecret)
	r := routes.SetupRouter(authService, cfg)
	log.Fatal(r.Run(":" + cfg.Port))
}
