package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string
	Text          *string
	Image         *string
	AuthorID      uint
	SubredditName string `gorm:"size:128"`
	Comments      []Comment
	Votes         []Vote
}
