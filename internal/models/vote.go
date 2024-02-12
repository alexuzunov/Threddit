package models

import "gorm.io/gorm"

type VoteType string

const (
	Upvote   VoteType = "upvote"
	Downvote VoteType = "downvote"
)

type Vote struct {
	gorm.Model
	Type      VoteType `gorm:"type:enum('upvote', 'downvote')"`
	AuthorID  uint
	PostID    *uint
	CommentID *uint
}
