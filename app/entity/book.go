package entity

import (
	"context"

	"github.com/working/go-clean-architecture/domain"
)

type BookEntity struct {
	bookRepo domain.BookRepository
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewBookEntity(a domain.BookRepository) domain.BookEntity {
	return &BookEntity{
		bookRepo: a,
	}
}
func (a *BookEntity) Fetch(c context.Context) (res []domain.Book, err error) {
	res, err = a.bookRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}

func (a *BookEntity) GetByID(c context.Context, id string) (res domain.Book, err error) {
	res, err = a.bookRepo.GetByID(c, id)
	return
}
