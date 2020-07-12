package content

import (
	"com.capturetweet/api"
	"context"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) api.ContentService {
	return &serviceImpl{repo}
}

func (s serviceImpl) SendMail(ctx context.Context, senderMail, senderName, message string) error {
	return s.repo.ContactUs(ctx, senderMail, senderMail, message)
}
