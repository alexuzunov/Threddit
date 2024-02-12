package repositories

import "gorm.io/gorm"

type UserRepository struct {
	*gorm.DB
}
