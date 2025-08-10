package viewsets

import (
	"go-crud/models"
	"go-crud/services"

	"github.com/gin-gonic/gin"
)

// PostViewSet extends BaseViewSet with Post-specific functionality
type PostViewSet struct {
	*BaseViewSet[models.Post]
	service *services.PostService
}

// NewPostViewSet creates a new PostViewSet instance
func NewPostViewSet() *PostViewSet {
	service := services.NewPostService()
	return &PostViewSet{
		BaseViewSet: NewBaseViewSet[models.Post](service),
		service:     service,
	}
}

// RegisterRoutes registers all Post-related routes
func (v *PostViewSet) RegisterRoutes(router *gin.Engine) {
	// Register base CRUD routes
	v.BaseViewSet.RegisterRoutes(router, "/posts")

	// You can add custom routes here if needed
	// For example:
	// postGroup := router.Group("/api/posts")
	// postGroup.GET("/published", v.ListPublished)
	// postGroup.POST("/:id/like", v.LikePost)
}

// Example of custom endpoint (commented out for now)
// You can uncomment and implement custom business logic

// ListPublished handles GET /api/posts/published
// func (v *PostViewSet) ListPublished(c *gin.Context) {
//     // Custom logic for published posts
//     posts, err := v.service.GetPublished()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, Response{
//             Error: fmt.Sprintf("Failed to fetch published posts: %v", err),
//         })
//         return
//     }
//
//     c.JSON(http.StatusOK, Response{
//         Data:  posts,
//         Count: len(posts),
//     })
// }

// LikePost handles POST /api/posts/:id/like
// func (v *PostViewSet) LikePost(c *gin.Context) {
//     id, err := strconv.ParseUint(c.Param("id"), 10, 32)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, Response{
//             Error: "Invalid ID format",
//         })
//         return
//     }
//
//     err = v.service.LikePost(uint(id))
//     if err != nil {
//         c.JSON(http.StatusNotFound, Response{
//             Error: fmt.Sprintf("Failed to like post: %v", err),
//         })
//         return
//     }
//
//     c.JSON(http.StatusOK, Response{
//         Message: "Post liked successfully",
//     })
// }
