package graph

import (
	"com.capturetweet/internal/convert"
	"time"
)

type Author struct {
	ID              string  `json:"id"`
	UserName        string  `json:"userName"`
	ScreenName      *string `json:"screenName"`
	Bio             *string `json:"bio"`
	ProfileImageURL *string `json:"profileImageURL"`
}

type Resource struct {
	ID        string  `json:"id"`
	URL       string  `json:"url"`
	MediaType *string `json:"mediaType"`
	Width     *int    `json:"width"`
	Height    *int    `json:"height"`
}

type Tweet struct {
	ID              string      `json:"id"`
	FullText        string      `json:"fullText"`
	PostedAt        *time.Time  `json:"postedAt"`
	AuthorID        string      `json:"author_id"`
	CaptureURL      *string     `json:"captureURL"`
	CaptureThumbURL *string     `json:"captureThumbURL"`
	FavoriteCount   *int        `json:"favoriteCount"`
	Lang            *string     `json:"lang"`
	RetweetCount    *int        `json:"retweetCount"`
	Resources       []*Resource `json:"resources"`
}

func (c *Tweet) Author() *Author {
	author, err := _userService.FindById(c.AuthorID)
	if err != nil {
		return nil
	}

	return &Author{
		ID:              author.ID,
		UserName:        author.UserName,
		ScreenName:      convert.String(author.ScreenName),
		Bio:             convert.String(author.Bio),
		ProfileImageURL: author.ProfileImage,
	}
}
