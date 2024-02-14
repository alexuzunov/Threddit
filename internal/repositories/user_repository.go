package repositories

import (
	"Threddit/internal/models"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	*gorm.DB
}

func (s *UserRepository) CreateUser(u *models.User) error {
	result := s.Create(u)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func (s *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	result := s.Where(&models.User{Username: username}).First(&user)

	return user, result.Error
}
