package main

import (
	"go-crud/initializers"
	"go-crud/viewsets"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Initialize ViewSets
	postViewSet := viewsets.NewPostViewSet()

	// Register routes using ViewSet pattern
	postViewSet.RegisterRoutes(router)

	// Add a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "go-crud-api",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
