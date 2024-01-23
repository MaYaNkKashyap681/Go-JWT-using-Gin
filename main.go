package main

import (
	"os"

	routes "github.com/jwtgo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Creating a New Server
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func (c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for api-1"})
	})

	router.GET("/api-2", func (c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for api-2"})
	})


	router.Run(":" + port)
}