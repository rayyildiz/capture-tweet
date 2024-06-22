//go:generate go run go.uber.org/mock/mockgen@latest -package=api -self_package=capturetweet.com/api -destination=service_mock.go . UserService,TweetService,SearchService,BrowserService,ContentService
package api

import (
	"context"

	"github.com/ChimeraCoder/anaconda"
)

type UserService interface {
	FindById(ctx context.Context, id string) (*UserModel, error)
	FindOrCreate(ctx context.Context, user *anaconda.User) (*UserModel, error)
}

type TweetService interface {
	FindById(ctx context.Context, id string) (*TweetModel, error)
	Store(ctx context.Context, url string) (string, error)
	Search(ctx context.Context, term string, size, start, page int) ([]TweetModel, error)
	SearchByUser(ctx context.Context, userId string) ([]TweetModel, error)
	UpdateLargeImage(ctx context.Context, id, captureUrl string) error
	UpdateThumbImage(ctx context.Context, id, captureUrl string) error
}

type SearchService interface {
	Search(ctx context.Context, term string, size int) ([]SearchModel, error)
	Put(ctx context.Context, tweetId, fullText, author string) error
}

type BrowserService interface {
	// CaptureURL capture and return a image (PNG)
	CaptureURL(ctx context.Context, model *CaptureRequestModel) ([]byte, error)

	// SaveCapture saves in a bucket
	SaveCapture(ctx context.Context, originalImage []byte, model *CaptureRequestModel) (*CaptureResponseModel, error)

	CaptureSaveUpdateDatabase(ctx context.Context, model *CaptureRequestModel) (*CaptureResponseModel, error)

	Close()
}

type ContentService interface {
	StoreContactRequest(ctx context.Context, senderMail, senderName, message, captcha string) error
}
