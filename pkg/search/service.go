package search

import (
	"context"

	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

type serviceImpl struct {
	index infra.IndexInterface
}

func NewService(index infra.IndexInterface) api.SearchService {
	return &serviceImpl{index}
}

func (s serviceImpl) Search(ctx context.Context, term string, size int) ([]api.SearchModel, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	res, err := s.index.Search(term, opt.HitsPerPage(size))
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	var list []api.SearchModel

	err = res.UnmarshalHits(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s serviceImpl) Put(ctx context.Context, tweetId, fullText, author string) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	span.SetAttributes(label.String("tweetId", tweetId))

	_, err := s.index.SaveObject(api.SearchModel{
		TweetID:  tweetId,
		FullText: fullText,
		Author:   author,
	})

	return err
}
