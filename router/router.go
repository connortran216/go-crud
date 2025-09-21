package router

import (
	"go-crud/initializers"
	"go-crud/views"

	"github.com/gin-gonic/gin"
)

// SetupRouter creates and configures the Gin router
func SetupRouter() *gin.Engine {
	// Initialize dependencies
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	router := gin.Default()

	postViews := views.NewPostViews()
	postViews.RegisterRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "go-crud-api",
		})
	})

	return router
}