package model

import (
	"time"
)

type Tweet struct {
	ID              string `gorm:"primary_key;type:varchar(36)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
	FullText        string     `gorm:"NOT NULL"`
	CaptureURL      string     `gorm:"NOT NULL"`
	CaptureThumbURL string     `gorm:"NOT NULL"`
	Lang            string
	FavoriteCount   int
	RetweetCount    int
	UserID          string `gorm:"type:varchar(36);NOT NULL"`
	User            User
}

func (Tweet) TableName() string { return "tweets" }

type User struct {
	ID         string `gorm:"primary_key;type:varchar(36)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserName   string     `gorm:"NOT NULL"`
	ScreenName string
}

func (User) TableName() string { return "users" }

type Resource struct {
	ID        string `gorm:"primary_key;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	URL       string     `gorm:"NOT NULL"`
	Width     int
	Height    int
	MediaType string
	TweetID   string `gorm:"type:varchar(36);NOT NULL"`
	Tweet     Tweet
}

func (Resource) TableName() string { return "resources" }
