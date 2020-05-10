package domain

import (
	"context"
)

// Book ...
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookUsecase interface {
	Fetch(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id string) (Book, error)
	// Update(ctx context.Context, ar *Article) error
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Store(context.Context, *Article) error
	// Delete(ctx context.Context, id int64) error
}

type BookRepository interface {
	Fetch(ctx context.Context) (res []Book, err error)
	GetByID(ctx context.Context, id string) (Book, error)
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Update(ctx context.Context, ar *Article) error
	// Store(ctx context.Context, a *Article) error
	// Delete(ctx context.Context, id int64) error
}
