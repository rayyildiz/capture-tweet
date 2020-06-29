//go:generate mockgen -package=service -self_package=com.capturetweet/pkg/service -destination=interfaces_mock.go . UserService,TweetService,SearchService,BrowserService,ContentService
package service

import (
	"github.com/ChimeraCoder/anaconda"
)

type UserService interface {
	FindById(id string) (*UserModel, error)
	FindOrCreate(user *anaconda.User) (*UserModel, error)
}

type TweetService interface {
	FindById(id string) (*TweetModel, error)
	Store(url string) (string, error)
	Search(term string, size, start, page int) ([]TweetModel, error)
	UpdateCaptureImage(id, captureUrl, captureThumbUrl string) error
}

type SearchService interface {
	Search(term string, size int) ([]SearchModel, error)
	Put(tweetId, fullText, author string) error
}

type BrowserService interface {
	// Capture and return a image (PNG)
	CaptureURL(model *CaptureRequestModel) ([]byte, error)

	// Save in a bucket
	SaveCapture(originalImage []byte, model *CaptureRequestModel) (*CaptureResponseModel, error)

	CaptureSaveUpdateDatabase(model *CaptureRequestModel) (*CaptureResponseModel, error)

	Close()
}

type ContentService interface {
	SendMail(senderMail, senderName, message string) error
}
