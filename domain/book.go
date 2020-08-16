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

type BookEntity interface {
	Fetch(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id string) (Book, error)
}

type BookRepository interface {
	Fetch(ctx context.Context) (res []Book, err error)
	GetByID(ctx context.Context, id string) (Book, error)
}
