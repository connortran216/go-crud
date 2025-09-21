package test

import (
	"bytes"
	"encoding/json"
	"go-crud/schemas"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreatePost_Success(t *testing.T) {
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

func TestCreatePost_ValidationError(t *testing.T) {
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