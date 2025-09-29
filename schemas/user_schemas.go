package schemas

import "go-crud/models"

type CreateUserInput struct {
	Name         string `json:"name" validate:"required" example:"Connor Tran"`
	Email        string `json:"email" validate:"required,email" example:"connor@example.com"`
	Password string `json:"password" validate:"required" example:"abcxyz123"`
}

type PartialUpdateUserInput struct {
	Name  *string `json:"name" example:"Connor Tran"`
	Email *string `json:"email" validate:"email" example:"connor@example.com"`
	Password *string `json:"password" example:"abcxyz123"`
}

type UserResponse struct {
	Data    models.User `json:"data"`
	Message string      `json:"message" example:"User created successfully"`
}
