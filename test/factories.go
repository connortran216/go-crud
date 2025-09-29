package test

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/brianvoe/gofakeit/v6"
)


type PostOption func(*models.Post)

func WithTitle(title string) PostOption {
    return func(p *models.Post) {
        p.Title = title
    }
}

func WithContent(content string) PostOption {
    return func(p *models.Post) {
        p.Content = content
    }
}

func PostFactory(opts ...PostOption) models.Post {
	post := &models.Post{
		Title:  gofakeit.Sentence(6),
		Content: gofakeit.Paragraph(1, 3, 12, " "),
	}

	for _, opt := range opts {
		opt(post)
	}

	initializers.DB.Create(post)
	return *post
}

type UserOption func(*models.User)

func WithEmail(email string) UserOption {
	return func(u *models.User) {
		u.Email = email
	}
}

func WithName(name string) UserOption {
	return func(u *models.User) {
		u.Name = name
	}
}

func UserFactory(opts ...UserOption) models.User {
	user := &models.User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		HashedPassword: "hashedPassword", 
	}

	for _, opt := range opts {
		opt(user)
	}

	initializers.DB.Create(user)
	return *user
}
