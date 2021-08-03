package internal

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(model MinerResult) (MinerResult, error) {
	err := r.db.Model(&model).Transaction(func(tx *gorm.DB) error {
		tx.Model(&model).Save(&model)

		return tx.Error
	})
	if err != nil {
		return MinerResult{}, err
	}

	return model, nil
}
