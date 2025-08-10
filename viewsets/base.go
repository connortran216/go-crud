package viewsets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count,omitempty"`
}

// ModelInterface defines the basic operations any model should support
type ModelInterface interface {
	GetID() uint
}

// ServiceInterface defines the service layer operations
type ServiceInterface[T ModelInterface] interface {
	Create(data T) (*T, error)
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(id uint, data T) (*T, error)
	PartialUpdate(id uint, partialData map[string]interface{}) (*T, error)
	Delete(id uint) error
}

// BaseViewSet provides generic CRUD operations
type BaseViewSet[T ModelInterface] struct {
	Service ServiceInterface[T]
}

// NewBaseViewSet creates a new base viewset
func NewBaseViewSet[T ModelInterface](service ServiceInterface[T]) *BaseViewSet[T] {
	return &BaseViewSet[T]{
		Service: service,
	}
}

// Create handles POST requests
func (v *BaseViewSet[T]) Create(c *gin.Context) {
	var data T
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	result, err := v.Service.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: fmt.Sprintf("Failed to create resource: %v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Data:    result,
		Message: "Resource created successfully",
	})
}

// List handles GET requests for multiple resources
func (v *BaseViewSet[T]) List(c *gin.Context) {
	results, err := v.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: fmt.Sprintf("Failed to fetch resources: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data:  results,
		Count: len(results),
	})
}

// Retrieve handles GET requests for a single resource
func (v *BaseViewSet[T]) Retrieve(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "Invalid ID format",
		})
		return
	}

	result, err := v.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Error: fmt.Sprintf("Resource not found: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: result,
	})
}

// Update handles PUT requests
func (v *BaseViewSet[T]) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "Invalid ID format",
		})
		return
	}

	var data T
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	result, err := v.Service.Update(uint(id), data)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Error: fmt.Sprintf("Failed to update resource: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data:    result,
		Message: "Resource updated successfully",
	})
}

// PartialUpdate handles PATCH requests for partial updates
func (v *BaseViewSet[T]) PartialUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "Invalid ID format",
		})
		return
	}

	// Parse partial data from request body
	var partialData map[string]interface{}
	if err := c.ShouldBindJSON(&partialData); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	// Check if there's any data to update
	if len(partialData) == 0 {
		c.JSON(http.StatusBadRequest, Response{
			Error: "No data provided for update",
		})
		return
	}

	// Update the resource with partial data
	result, err := v.Service.PartialUpdate(uint(id), partialData)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "post not found" || err.Error() == "resource not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "title cannot be empty" || err.Error() == "content cannot be empty" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, Response{
			Error: fmt.Sprintf("Failed to update resource: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data:    result,
		Message: "Resource updated successfully",
	})
}

// Destroy handles DELETE requests
func (v *BaseViewSet[T]) Destroy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "Invalid ID format",
		})
		return
	}

	err = v.Service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Error: fmt.Sprintf("Failed to delete resource: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: "Resource deleted successfully",
	})
}

// RegisterRoutes automatically registers CRUD routes for the viewset
func (v *BaseViewSet[T]) RegisterRoutes(router *gin.Engine, basePath string) {
	group := router.Group(basePath)
	{
		group.POST("", v.Create)
		group.GET("", v.List)
		group.GET("/:id", v.Retrieve)
		group.PUT("/:id", v.Update)
		group.PATCH("/:id", v.PartialUpdate)
		group.DELETE("/:id", v.Destroy)
	}
}
