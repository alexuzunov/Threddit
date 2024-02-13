package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	Admin    Role = "admin"
	Redditor Role = "redditor"
)

type User struct {
	gorm.Model
	Email         string       `json:"email,omitempty"`
	Username      string       `gorm:"unique" json:"username,omitempty"`
	Password      string       `json:"password,omitempty"`
	Image         string       `json:"image,omitempty"`
	Followers     []*User      `gorm:"many2many:followers"`
	Subreddits    []Subreddit  `gorm:"foreignkey:CreatorID"`
	Posts         []Post       `gorm:"foreignkey:AuthorID"`
	Comments      []Comment    `gorm:"foreignkey:AuthorID"`
	Subscriptions []*Subreddit `gorm:"many2many:subscribers"`
	Votes         []Vote       `gorm:"foreignkey:AuthorID"`
	Role          Role         `json:"role,omitempty"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}
