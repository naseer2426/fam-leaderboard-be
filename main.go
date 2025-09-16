package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naseer2426/fam-leaderboard-be/internal/api"
	db "github.com/naseer2426/fam-leaderboard-be/internal/db"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		// Only log if the file is missing; envs may be provided by the environment
		if !os.IsNotExist(err) {
			log.Printf("warning: could not load .env: %v", err)
		}
	}
	// Initialize database and run migrations
	db.AutoMigrate()

	router := gin.Default()
	router.GET("/", api.HelloWorld)
	router.GET("/scores", api.GetScores)
	router.POST("/scores", api.CreateScore)
	router.PUT("/scores/increase_score", api.IncreaseScore)
	router.PUT("/scores/decrease_score", api.DecreaseScore)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
