package test


import (
	"go-crud/initializers"
	"go-crud/models"
	"go-crud/router"
	"testing"

	"github.com/gin-gonic/gin"
)


type BaseTestSuite struct {
	router *gin.Engine
	t *testing.T
}

func NewTestSuite(t *testing.T) *BaseTestSuite {
	suite := &BaseTestSuite{t: t}
	suite.SetUp()
	return suite
}

func (suite *BaseTestSuite) SetUp() {
	gin.SetMode(gin.TestMode)
	suite.router = router.SetupRouter()
	suite.CleanUp()
}

func (suite *BaseTestSuite) CleanUp() {
	initializers.DB.Where("title LIKE ?", "Test%").Delete(&models.Post{})
}

func (suite *BaseTestSuite) TearDown() {
	suite.CleanUp()
}

