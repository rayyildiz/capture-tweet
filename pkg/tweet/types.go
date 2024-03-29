package tweet

import (
	"time"
)

type Tweet struct {
	ID               string     `docstore:"id"`
	CreatedAt        time.Time  `docstore:"created_at"`
	UpdatedAt        time.Time  `docstore:"updated_at"`
	PostedAt         time.Time  `docstore:"posted_at"`
	FullText         string     `docstore:"full_text"`
	CaptureURL       *string    `docstore:"capture_url"`
	CaptureThumbURL  *string    `docstore:"capture_thumb_url"`
	Lang             string     `docstore:"lang"`
	FavoriteCount    int        `docstore:"favorite_count"`
	RetweetCount     int        `docstore:"retweet_count"`
	AuthorID         string     `docstore:"author_id"`
	Resources        []Resource `docstore:"resources"`
	DocstoreRevision interface{}
}

type SortByPosted []Tweet

func (a SortByPosted) Len() int           { return len(a) }
func (a SortByPosted) Less(i, j int) bool { return a[i].PostedAt.After(a[j].PostedAt) }
func (a SortByPosted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Resource struct {
	ID        string `docstore:"id"`
	URL       string `docstore:"url"`
	Width     int    `docstore:"width"`
	Height    int    `docstore:"height"`
	MediaType string `docstore:"media_type"`
}
