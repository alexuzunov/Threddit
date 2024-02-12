package models

import (
	"gorm.io/gorm"
	"time"
)

type Subreddit struct {
	Name        string `gorm:"primaryKey" gorm:"size:128"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string
	Subscribers []*User `gorm:"many2many:subscribers"`
	Posts       []Post
	CreatorID   uint
}
