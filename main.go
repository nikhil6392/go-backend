package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhil6392/go-backend/config"
	"github.com/nikhil6392/go-backend/routes"
)



func main() {

	config.InitDB()

	router := gin.Default()

	routes.Register(router)
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "Welcome to the golang backend"})
	})

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/db-status", func(c *gin.Context){
		var now time.Time
		err := config.DB.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"db_time": now})
	})

	// Start server on port 8080
	router.Run(":8080")
}
