package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}


func ListPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}


func GetPostById(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	// Get post by id
	initializers.DB.First(&post, c.Param("id"))
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save post
	initializers.DB.Save(&post)

	// Return post
	c.JSON(http.StatusOK, gin.H{"data": post})
}


func DeletePost(c *gin.Context) {
	var post models.Post
	initializers.DB.Delete(&post, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}


