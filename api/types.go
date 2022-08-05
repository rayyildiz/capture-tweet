package api

import (
	"time"
)

type UserModel struct {
	ProfileImage *string
	ID           string
	UserName     string
	ScreenName   string
	Bio          string
}

type TweetModel struct {
	PostedAt        *time.Time
	CaptureURL      *string
	CaptureThumbURL *string
	Author          *UserModel
	ID              string
	FullText        string
	Lang            string
	AuthorID        string
	Resources       []ResourceModel
	FavoriteCount   int
	RetweetCount    int
}

type ResourceModel struct {
	ID           string
	URL          string
	ResourceType string
	Width        int
	Height       int
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

type CaptureResponseModel struct {
	ID         string
	CaptureURL string
}
