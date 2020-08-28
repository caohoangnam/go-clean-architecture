package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/working/go-clean-architecture/domain"
)

type MeowRepository struct {
	Conn *gorm.DB
}

func NewMeowRepository(Conn *gorm.DB) domain.MeowRepository {
	return &MeowRepository{Conn}
}

func (m MeowRepository) Close() {
	m.Conn.Close()
}

func (m MeowRepository) Create(ctx context.Context, meow domain.Meow) error {
	fmt.Println("meow", meow)
	err := m.Conn.Create(&meow).Error
	return err
}

func (m MeowRepository) List(ctx context.Context, skip, take int64) ([]domain.Meow, error) {
	rows, err := m.Conn.Raw("SELECT * FROM meows ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take).Rows()
	if err != nil {
		return nil, err
	}

	meows := []domain.Meow{}
	for rows.Next() {
		meow := domain.Meow{}
		if err = rows.Scan(&meow.Id, &meow.Body, &meow.CreatedAt); err == nil {
			meows = append(meows, meow)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return meows, nil
}
