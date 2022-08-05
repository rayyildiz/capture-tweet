package resolver

import (
	"context"
	"time"

	"capturetweet.com/internal/convert"
)

type Author struct {
	ScreenName      *string `json:"screenName"`
	Bio             *string `json:"bio"`
	ProfileImageURL *string `json:"profileImageURL"`
	ID              string  `json:"id"`
	UserName        string  `json:"userName"`
}

type Resource struct {
	MediaType *string `json:"mediaType"`
	Width     *int    `json:"width"`
	Height    *int    `json:"height"`
	ID        string  `json:"id"`
	URL       string  `json:"url"`
}

type Tweet struct {
	ID              string      `json:"id"`
	FullText        string      `json:"fullText"`
	PostedAt        *time.Time  `json:"postedAt"`
	AuthorID        *string     `json:"authorID"`
	CaptureURL      *string     `json:"captureURL"`
	CaptureThumbURL *string     `json:"captureThumbURL"`
	FavoriteCount   *int        `json:"favoriteCount"`
	Lang            *string     `json:"lang"`
	RetweetCount    *int        `json:"retweetCount"`
	Resources       []*Resource `json:"resources"`
}

func (c *Tweet) Author() *Author {
	if c.AuthorID == nil {
		return nil
	}
	author, err := _userService.FindById(context.Background(), *c.AuthorID)
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
