package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/working/go-clean-architecture/domain"
)

type BookRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewBookRepository(Conn *gorm.DB) domain.BookRepository {
	return &BookRepository{Conn}
}

func (m *BookRepository) Fetch(ctx context.Context) (res []domain.Book, err error) {
	var books []domain.Book
	m.Conn.Find(&books)

	return books, nil
}
func (m *BookRepository) GetByID(ctx context.Context, id string) (res domain.Book, err error) {
	var book domain.Book
	m.Conn.Where("id = ?", id).First(&book)

	return book, nil
}
