package content

import (
	"com.capturetweet/api"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) api.ContentService {
	return &serviceImpl{repo}
}

func (s serviceImpl) SendMail(senderMail, senderName, message string) error {
	return s.repo.ContactUs(senderMail, senderMail, message)
}
