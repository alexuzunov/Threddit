package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	*UserRepository
	*SubredditRepository
	*PostRepository
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	return &Repository{
		UserRepository:      &UserRepository{DB: db},
		SubredditRepository: &SubredditRepository{DB: db},
		PostRepository:      &PostRepository{DB: db},
	}, nil
}
