package tweet

import (
	"time"
)

type Tweet struct {
	ID              string     `db:"id"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	PostedAt        time.Time  `db:"posted_at"`
	FullText        string     `db:"full_text"`
	CaptureURL      *string    `db:"capture_url"`
	CaptureThumbURL *string    `db:"capture_thumb_url"`
	Lang            string     `db:"lang"`
	FavoriteCount   int        `db:"favorite_count"`
	RetweetCount    int        `db:"retweet_count"`
	AuthorID        string     `db:"author_id"`
	Resources       []Resource `db:"resources"`
}

type SortByPosted []Tweet

func (a SortByPosted) Len() int           { return len(a) }
func (a SortByPosted) Less(i, j int) bool { return a[i].PostedAt.After(a[j].PostedAt) }
func (a SortByPosted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Resource struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MediaType string `json:"media_type"`
}
