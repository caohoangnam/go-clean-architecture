package domain

import (
	"context"
	"time"
)

type Meow struct {
	Id        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type MeowEntity interface {
	Create(ctx context.Context, meo Meow) error
	List(ctx context.Context, skip, take int64) ([]Meow, error)
	SearchMeows(ctx context.Context, search string, skip, take int64) ([]Meow, error)
}

type MeowRepository interface {
	Create(ctx context.Context, meo Meow) error
	List(ctx context.Context, skip, take int64) ([]Meow, error)
	SearchMeows(ctx context.Context, search string, skip, take int64) ([]Meow, error)
}
