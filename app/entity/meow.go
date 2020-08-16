package entity

import (
	"context"

	"github.com/working/go-clean-architecture/domain"
)

type MeowEntity struct {
	meowRepo domain.MeowRepository
}

func NewMeowEntity(meow domain.MeowRepository) domain.MeowEntity {
	return &MeowEntity{
		meowRepo: meow,
	}
}

func (m *MeowEntity) Create(c context.Context, meow domain.Meow) error {
	err := m.meowRepo.Create(c, meow)
	if err != nil {
		return err
	}
	return nil
}

func (m *MeowEntity) List(c context.Context, skip, take int64) ([]domain.Meow, error) {
	meows, err := m.meowRepo.List(c, skip, take)
	if err != nil {
		return nil, err
	}
	return meows, nil
}
