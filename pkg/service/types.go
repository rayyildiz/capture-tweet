package service

type UserModel struct {
	ID         string
	UserName   string
	ScreenName string
}

type TweetModel struct {
	ID            string
	FullText      string
	Lang          string
	CaptureURL    *string
	ThumbnailURL  *string
	FavoriteCount int
	RetweetCount  int
	Author        *UserModel
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
