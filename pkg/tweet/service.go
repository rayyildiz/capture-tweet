package tweet

import (
	"com.capturetweet/pkg/service"
)

type serviceImpl struct {
	repo   Repository
	search service.SearchService
}

func NewService(repo Repository, search service.SearchService) service.TweetService {
	return &serviceImpl{repo, search}
}

func (s serviceImpl) FindById(id string) (*service.TweetModel, error) {
	return nil, nil
}

func (s serviceImpl) Store(tweet *service.TweetModel, user *service.UserModel, resources []service.ResourceModel) error {
	err := s.repo.Store(tweet.ID, tweet.FullText, tweet.Lang, tweet.Author.ID, tweet.RetweetCount, tweet.FavoriteCount, nil)
	if err != nil {
		return err
	}
	go func(t *service.TweetModel) {
		s.search.Put(t.ID, t.FullText, t.Author.UserName)
	}(tweet)

	return nil
}

func (s serviceImpl) Search(term string, size, start, page int) ([]service.TweetModel, error) {
	searchModels, err := s.search.Search(term, size)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, model := range searchModels {
		ids = append(ids, model.TweetID)
	}

	tweets, err := s.repo.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	var res []service.TweetModel

	for _, tweet := range tweets {
		res = append(res, service.TweetModel{
			ID:            tweet.ID,
			FullText:      tweet.FullText,
			Lang:          tweet.Lang,
			CaptureURL:    nil,
			ThumbnailURL:  nil,
			FavoriteCount: tweet.FavoriteCount,
			RetweetCount:  tweet.RetweetCount,
			Author:        nil,
		})
	}

	return res, nil
}
