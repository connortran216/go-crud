package services

import (
	"errors"
	"go-crud/initializers"
	"go-crud/models"
	"go-crud/schemas"

	"gorm.io/gorm"
)

// PostService handles business logic for Post operations
type PostService struct {
	db *gorm.DB
}

// NewPostService creates a new PostService instance
func NewPostService() *PostService {
	return &PostService{
		db: initializers.DB,
	}
}

// Create creates a new post
func (s *PostService) Create(post models.Post) (*models.Post, error) {
	if post.Title == "" {
		return nil, errors.New("title is required")
	}
	if post.Content == "" {
		return nil, errors.New("content is required")
	}

	result := s.db.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

// GetByID retrieves a post by ID
func (s *PostService) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	result := s.db.First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, result.Error
	}

	return &post, nil
}

// GetAll retrieves all posts
func (s *PostService) GetAll() ([]models.Post, error) {
	var posts []models.Post
	result := s.db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

// GetPaginated retrieves posts with pagination
func (s *PostService) GetWithPagination(query schemas.ListPostsQueryParams) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64
	
	// Get total count
	if err := s.db.Model(&models.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Calculate offset
	offset := (query.Page - 1) * query.Limit
	
	// Get paginated results
	result := s.db.Limit(query.Limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	
	return posts, total, nil
}

// Update updates an existing post
func (s *PostService) Update(id uint, updatedPost models.Post) (*models.Post, error) {
	var post models.Post

	// Check if post exists
	result := s.db.First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, result.Error
	}

	// Validate updated data
	if updatedPost.Title == "" {
		return nil, errors.New("title is required")
	}
	if updatedPost.Content == "" {
		return nil, errors.New("content is required")
	}

	// Update fields
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	// Save changes
	result = s.db.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

// PartialUpdate updates specific fields of an existing post
func (s *PostService) PartialUpdate(id uint, partialData map[string]interface{}) (*models.Post, error) {
	var post models.Post

	// Check if post exists
	result := s.db.First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, result.Error
	}

	// Update only provided fields
	if title, exists := partialData["title"]; exists {
		if titleStr, ok := title.(string); ok && titleStr != "" {
			post.Title = titleStr
		} else if titleStr == "" {
			return nil, errors.New("title cannot be empty")
		}
	}

	if content, exists := partialData["content"]; exists {
		if contentStr, ok := content.(string); ok && contentStr != "" {
			post.Content = contentStr
		} else if contentStr == "" {
			return nil, errors.New("content cannot be empty")
		}
	}

	// Save changes
	result = s.db.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

// Delete deletes a post by ID
func (s *PostService) Delete(id uint) error {
	var post models.Post

	// Check if post exists
	result := s.db.First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return result.Error
	}

	// Delete the post
	result = s.db.Delete(&post)
	return result.Error
}
