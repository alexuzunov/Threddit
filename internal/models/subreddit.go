package models

import (
	"gorm.io/gorm"
	"time"
)

type SubredditType string

const (
	Public     SubredditType = "public"
	Restricted SubredditType = "restricted"
	Private    SubredditType = "private"
)

type Subreddit struct {
	Name        string `gorm:"primaryKey" gorm:"size:128"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string
	Image       string
	Type        SubredditType `gorm:"type:enum('public', 'restricted', 'private')"`
	NSFW        bool
	Subscribers []*User `gorm:"many2many:subscribers"`
	Posts       []Post
	CreatorID   uint
}
