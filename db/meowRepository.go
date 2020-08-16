package db

import (
	"context"

	"github.com/working/go-clean-architecture/domain"
)

type MeowReposiory interface {
	Close()
	Create(ctx context.Context, meow domain.Meow) error
	List(ctx context.Context, skip, take int64) ([]domain.Meow, error)
}

var impl MeowRepository

func SetRepository(repo MeowRepository) {
	impl = repo
}

func Close() {
	impl.Close()
}

func Create(ctx context.Context, meow domain.Meow) error {
	return impl.Create(ctx, meow)
}

func List(ctx context.Context, skip, take int64) ([]domain.Meow, error) {
	return impl.List(ctx, skip, take)
}
