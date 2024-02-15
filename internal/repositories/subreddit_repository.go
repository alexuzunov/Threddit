package repositories

import (
	"Threddit/internal/models"
	"gorm.io/gorm"
	"log"
)

type SubredditRepository struct {
	*gorm.DB
}

func (s *SubredditRepository) CreateSubreddit(sr *models.Subreddit) error {
	result := s.Create(sr)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func (s *SubredditRepository) GetCreatedSubreddits(username string) ([]models.Subreddit, error) {
	var user models.User
	result := s.Preload("Subreddits").Where(&models.User{Username: username}).First(&user)

	return user.Subreddits, result.Error
}
