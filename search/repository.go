package search

import (
	"context"

	"github.com/working/go-clean-architecture/domain"
)

var impl Repository

type Repository interface {
	Close()
	Create(ctx context.Context, meow domain.Meow) error
	Search(ctx context.Context, query string, skip int64, take int64) ([]domain.Meow, error)
}

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func Create(ctx context.Context, meow domain.Meow) error {
	return impl.Create(ctx, meow)
}

func Search(ctx context.Context, query string, skip int64, take int64) ([]domain.Meow, error) {
	return impl.Search(ctx, query, skip, take)
}
