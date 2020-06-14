package tweet

import (
	"com.capturetweet/pkg/service"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) service.TweetService {
	return &serviceImpl{repo}
}

func (s serviceImpl) FindById(id string) (*service.TweetModel, error) {
	return nil, nil
}

func (s serviceImpl) Store(tweet *service.TweetModel, user *service.UserModel, resources []service.ResourceModel) error {
	return nil
}

func (s serviceImpl) Search(term string, size int, cursorId *string) ([]service.TweetModel, error) {
	return nil, nil
}
