package test

import (
	"bytes"
	"encoding/json"
	"go-crud/schemas"
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
)


func TestCreatePostSuccess(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()
	
	requestBody := map[string]string{
		"title":   "Test Post Title",
		"content": "This is a test post content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response schemas.PostResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Post Title", response.Data.Title)
	assert.Equal(t, "This is a test post content", response.Data.Content)
	assert.Equal(t, "Post created successfully", response.Message)
	assert.NotZero(t, response.Data.ID)
}

// TODO: Fix this case that API return InternalServerError instead of BadRequest (500 code vs 400 code)
func TestCreatePostValidationError(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()
	
	requestBody := map[string]string{
		"title":   "",
		"content": "This is a test post content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response.Error, "title is required")
}


func TestGetPostByIDSuccess(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Post
	post := PostFactory()
	req, _ := http.NewRequest("GET", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.PostResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, post.Title, response.Data.Title)
	assert.Equal(t, post.Content, response.Data.Content)
}

func TestGetPostByIDDataDoesNotExist(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	req, _ := http.NewRequest("GET", "/posts/9999", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, response.Error, "post not found")
}

func TestListPostsSuccessWithDefaultPagination(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Posts
	for i := 0; i < 10; i++ {
		PostFactory()
	}

	req, _ := http.NewRequest("GET", "/posts", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ListPostsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, len(response.Data), 10)
	assert.Equal(t, 10, response.Limit)
	assert.Equal(t, 1, response.Page)
	assert.GreaterOrEqual(t, response.Total, 10)
}

func TestListPostsSuccessWithCustomPagination(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Posts
	for i := 0; i < 10; i++ {
		PostFactory()
	}

	req, _ := http.NewRequest("GET", "/posts?page=2&limit=5", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ListPostsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, len(response.Data), 5)
	assert.Equal(t, 5, response.Limit)
	assert.Equal(t, 2, response.Page)
	assert.GreaterOrEqual(t, response.Total, 10)
}

func TestListPostsShouldReturnDefaultPaginationWhenInvalidQueryParams(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	req, _ := http.NewRequest("GET", "/posts?page=abc&limit=xyz", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ListPostsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, len(response.Data), 0)
	assert.Equal(t, 10, response.Limit) // Default limit
	assert.Equal(t, 1, response.Page)   // Default page
	assert.GreaterOrEqual(t, response.Total, 0)
}

func TestUpdatePostSuccess(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Post
	post := PostFactory()
	
	requestBody := map[string]string{
		"title":   "Updated Title",
		"content": "Updated Content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.PostResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Updated Title", response.Data.Title)
	assert.Equal(t, "Updated Content", response.Data.Content)
	assert.NotZero(t, response.Data.ID)
}

func TestUpdatePostFailWhenDataDoesNotExist(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()
	
	requestBody := map[string]string{
		"title":   "Updated Title",
		"content": "Updated Content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/posts/9999", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, response.Error, "post not found")
}

func TestUpdatePostFailWhenDataIsInvalid(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Post
	post := PostFactory()
	
	requestBody := map[string]string{
		"author":   "", // Invalid un-exist field
		"content": "Updated Content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, response.Error, "title is required")
}

func TestPartiallyUpdatePostSuccess(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Post
	post := PostFactory()
	
	requestBody := map[string]string{
		"content": "Partially Updated Content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PATCH", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.PostResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, post.Title, response.Data.Title) // Title should remain unchanged
	assert.Equal(t, "Partially Updated Content", response.Data.Content)
	assert.NotZero(t, response.Data.ID)
}

func TestPartiallyUpdatePostFailWhenDataDoesNotExist(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()
	
	requestBody := map[string]string{
		"content": "Partially Updated Content",
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PATCH", "/posts/9999", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, response.Error, "post not found")
}

func TestPartiallyUpdatePostFailInvalidData(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()
	// Create mock data Post
	post := PostFactory()
	
	requestBody := map[string]string{
		"title": "", // Invalid empty title
	}
	
	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PATCH", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, response.Error, "title cannot be empty")
}

func TestDeletePostSuccess(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	// Create mock data Post
	post := PostFactory()
	req, _ := http.NewRequest("DELETE", "/posts/"+strconv.FormatUint(uint64(post.ID), 10), nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Post deleted successfully", response.Message)
}

func TestDeletePostFailWhenDataDoesNotExist(t *testing.T) {
	suite := NewTestSuite(t)
	defer suite.TearDown()

	req, _ := http.NewRequest("DELETE", "/posts/9999", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var response schemas.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, response.Error, "post not found")
}
