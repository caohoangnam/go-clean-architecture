package psql

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/working/project/domain"
)

type psqlBookRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPsqlBookRepository(Conn *gorm.DB) domain.BookRepository {
	return &psqlBookRepository{Conn}
}

func (m *psqlBookRepository) Fetch(ctx context.Context) (res []domain.Book, err error) {
	var books []domain.Book
	m.Conn.Find(&books)

	return books, nil
}
func (m *psqlBookRepository) GetByID(ctx context.Context, id string) (res domain.Book, err error) {
	var book domain.Book
	m.Conn.Where("id = ?", id).First(&book)
	// if err := m.Conn.Where("id = ?", id).First(&book).Error; err != nil {
	// 	return err, nil
	// }

	return book, nil
}
