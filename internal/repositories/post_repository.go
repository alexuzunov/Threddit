package repositories

import (
	"Threddit/internal/models"
	"gorm.io/gorm"
	"log"
)

type PostRepository struct {
	*gorm.DB
}

func (s *PostRepository) CreatePost(p *models.Post) error {
	result := s.Create(p)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}
