package usecase

import (
	"context"

	"github.com/working/project/domain"
)

type bookUsecase struct {
	bookRepo domain.BookRepository
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewBookUsecase(a domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepo: a,
	}
}
func (a *bookUsecase) Fetch(c context.Context) (res []domain.Book, err error) {
	res, err = a.bookRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}

func (a *bookUsecase) GetByID(c context.Context, id string) (res domain.Book, err error) {
	res, err = a.bookRepo.GetByID(c, id)
	// if err != nil {
	// 	return nil, err
	// }
	return
}
