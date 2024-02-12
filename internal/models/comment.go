package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text      string
	AuthorID  uint
	PostID    uint
	ReplyToID *uint
	ReplyTo   *Comment
	CommentID *uint
	Replies   []Comment `gorm:"foreignkey:CommentID"`
	Votes     []Vote
}
