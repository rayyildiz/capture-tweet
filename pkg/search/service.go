package search

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/service"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type serviceImpl struct {
	index infra.IndexInterface
}

func NewService(index infra.IndexInterface) service.SearchService {
	return &serviceImpl{index}
}

func (s serviceImpl) Search(term string, size int) ([]service.SearchModel, error) {
	res, err := s.index.Search(term, opt.HitsPerPage(size))
	if err != nil {
		return nil, err
	}
	var list []service.SearchModel

	err = res.UnmarshalHits(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s serviceImpl) Put(tweetId, fullText, author string) error {
	_, err := s.index.SaveObject(service.SearchModel{
		TweetID:  tweetId,
		FullText: fullText,
		Author:   author,
	})

	return err
}
