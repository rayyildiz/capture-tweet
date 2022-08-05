package tweet

import (
	"time"
)

type Tweet struct {
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	PostedAt        time.Time  `db:"posted_at"`
	CaptureThumbURL *string    `db:"capture_thumb_url"`
	CaptureURL      *string    `db:"capture_url"`
	FullText        string     `db:"full_text"`
	ID              string     `db:"id"`
	Lang            string     `db:"lang"`
	AuthorID        string     `db:"author_id"`
	Resources       []Resource `db:"resources"`
	FavoriteCount   int        `db:"favorite_count"`
	RetweetCount    int        `db:"retweet_count"`
}

type SortByPosted []Tweet

func (a SortByPosted) Len() int           { return len(a) }
func (a SortByPosted) Less(i, j int) bool { return a[i].PostedAt.After(a[j].PostedAt) }
func (a SortByPosted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Resource struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	MediaType string `json:"media_type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
