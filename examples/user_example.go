package examples

// This is an example of how to create additional ViewSets for other models
// following the same pattern as PostViewSet

/*

// 1. First, create your model (in models/userModel.go)
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Implement ModelInterface
func (u *User) GetID() uint {
	return u.ID
}

// 2. Create service layer (in services/user_service.go)
type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{db: initializers.DB}
}

func (s *UserService) Create(user models.User) (*models.User, error) {
	// Business logic and validation
	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	result := s.db.Create(&user)
	return &user, result.Error
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	return users, result.Error
}

func (s *UserService) Update(id uint, updatedUser models.User) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	result := s.db.Save(&user)
	return &user, result.Error
}

func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

// 3. Create ViewSet (in viewsets/user_viewset.go)
type UserViewSet struct {
	*BaseViewSet[models.User]
	service *services.UserService
}

func NewUserViewSet() *UserViewSet {
	service := services.NewUserService()
	return &UserViewSet{
		BaseViewSet: NewBaseViewSet[models.User](service),
		service:     service,
	}
}

func (v *UserViewSet) RegisterRoutes(router *gin.Engine) {
	v.BaseViewSet.RegisterRoutes(router, "/api/users")

	// Add custom routes if needed
	userGroup := router.Group("/api/users")
	userGroup.GET("/by-email/:email", v.GetByEmail)
}

func (v *UserViewSet) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	// Custom business logic here
}

// 4. Register in main.go
func main() {
	router := gin.Default()

	// Initialize ViewSets
	postViewSet := viewsets.NewPostViewSet()
	userViewSet := viewsets.NewUserViewSet()

	// Register routes
	postViewSet.RegisterRoutes(router)
	userViewSet.RegisterRoutes(router)

	router.Run()
}

*/
