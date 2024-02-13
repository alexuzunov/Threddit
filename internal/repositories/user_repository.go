package repositories

import (
	"Threddit/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func (s *UserRepository) CreateUser(u *models.User) error {
	return nil
}
