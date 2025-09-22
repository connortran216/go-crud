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
