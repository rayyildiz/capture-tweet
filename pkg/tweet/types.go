package tweet

import (
	"time"
)

type Tweet struct {
	ID               string    `docstore:"id"`
	CreatedAt        time.Time `docstore:"created_at"`
	UpdatedAt        time.Time `docstore:"updated_at"`
	PostedAt         time.Time `docstore:"posted_at"`
	FullText         string
	CaptureURL       *string
	CaptureThumbURL  *string
	Lang             string
	FavoriteCount    int
	RetweetCount     int
	UserID           string
	Resources        []Resource
	DocstoreRevision interface{}
}

type Resource struct {
	ID        string
	CreatedAt time.Time
	URL       string
	Width     int
	Height    int
	MediaType string
}
