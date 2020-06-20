package service

import (
	"time"
)

type UserModel struct {
	ID           string
	UserName     string
	ScreenName   string
	Bio          string
	ProfileImage *string
}

type TweetModel struct {
	ID              string
	FullText        string
	Lang            string
	PostedAt        *time.Time
	CaptureURL      *string
	CaptureThumbURL *string
	FavoriteCount   int
	RetweetCount    int
	Author          *UserModel
}

type ResourceModel struct {
	ID           string
	URL          string
	Width        int
	Height       int
	ResourceType string
}

type SearchModel struct {
	TweetID  string `json:"tweet_id"`
	FullText string `json:"full_text"`
	Author   string `json:"author"`
}
