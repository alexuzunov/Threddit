package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	*UserRepository
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	return &Repository{
		UserRepository: &UserRepository{DB: db},
	}, nil
}
