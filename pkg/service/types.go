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
	AuthorID        string
	Author          *UserModel
	Resources       []ResourceModel
}

type ResourceModel struct {
	ID           string
	URL          string
	Width        int
	Height       int
	ResourceType string
}

type SearchModel struct {
	TweetID  string `json:"objectID"`
	FullText string `json:"fullText"`
	Author   string `json:"author"`
}

type CaptureRequestModel struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Url    string `json:"url"`
}
