package views

import (
	"fmt"
	"go-crud/schemas"
	"go-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostViews struct {
	service *services.PostService
}

func NewPostViews() *PostViews {
	return &PostViews{
		service: services.NewPostService(),
	}
}

func (v *PostViews) CreatePost(c *gin.Context) {
	var input schemas.CreatePostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	result, err := v.service.Create(input.ToModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Error: fmt.Sprintf("Failed to create post: %v", err),
		})
		return
	}

	response := schemas.PostResponse{
		Data:    *result,
		Message: "Post created successfully",
	}
	c.JSON(http.StatusCreated, response)
}

func (v *PostViews) ListPosts(c *gin.Context) {
	var query schemas.ListPostsQueryParams
	query.Page, _ = strconv.Atoi(c.Query("page"))
	query.Limit, _ = strconv.Atoi(c.Query("limit"))
	
	// Set defaults if not provided
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Limit == 0 {
		query.Limit = 10
	}

	results, total, err := v.service.GetWithPagination(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Error: fmt.Sprintf("Failed to fetch posts: %v", err),
		})
		return
	}

	response := schemas.ListPostsResponse{
		Data:  results,
		Limit: query.Limit,
		Page:  query.Page,
		Total: int(total),
	}
	c.JSON(http.StatusOK, response)
}

func (v *PostViews) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: "Invalid ID format",
		})
		return
	}

	result, err := v.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Error: fmt.Sprintf("Post not found: %v", err),
		})
		return
	}

	response := schemas.PostResponse{
		Data: *result,
	}
	c.JSON(http.StatusOK, response)
}

func (v *PostViews) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: "Invalid ID format",
		})
		return
	}

	var input schemas.UpdatePostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	result, err := v.service.Update(uint(id), input.ToModel())
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Error: fmt.Sprintf("Failed to update post: %v", err),
		})
		return
	}

	response := schemas.PostResponse{
		Data:    *result,
		Message: "Post updated successfully",
	}
	c.JSON(http.StatusOK, response)
}

func (v *PostViews) PartialUpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: "Invalid ID format",
		})
		return
	}

	var input schemas.PatchPostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: fmt.Sprintf("Invalid request data: %v", err),
		})
		return
	}

	if input.IsEmpty() {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: "No data provided for update",
		})
		return
	}

	result, err := v.service.PartialUpdate(uint(id), input.ToMap())
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "post not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "title cannot be empty" || err.Error() == "content cannot be empty" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, schemas.ErrorResponse{
			Error: fmt.Sprintf("Failed to update post: %v", err),
		})
		return
	}

	response := schemas.PostResponse{
		Data:    *result,
		Message: "Post updated successfully",
	}
	c.JSON(http.StatusOK, response)
}

func (v *PostViews) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Error: "Invalid ID format",
		})
		return
	}

	err = v.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Error: fmt.Sprintf("Failed to delete post: %v", err),
		})
		return
	}

	response := schemas.MessageResponse{
		Message: "Post deleted successfully",
	}
	c.JSON(http.StatusOK, response)
}

func (v *PostViews) RegisterRoutes(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.POST("", v.CreatePost)
		posts.GET("", v.ListPosts)
		posts.GET("/:id", v.GetPost)
		posts.PUT("/:id", v.UpdatePost)
		posts.PATCH("/:id", v.PartialUpdatePost)
		posts.DELETE("/:id", v.DeletePost)
	}
}