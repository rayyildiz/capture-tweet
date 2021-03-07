package search

import (
	"context"
	"go.opentelemetry.io/otel"

	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

type serviceImpl struct {
	index  infra.IndexInterface
	tracer trace.Tracer
}

func NewService(index infra.IndexInterface) api.SearchService {
	return &serviceImpl{
		index:  index,
		tracer: otel.GetTracerProvider().Tracer("com.capturetweet/pkg/search"),
	}
}

func (s serviceImpl) Search(ctx context.Context, term string, size int) ([]api.SearchModel, error) {
	ctx, span := s.tracer.Start(ctx, "service:search")
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
	ctx, span := s.tracer.Start(ctx, "service:put")
	defer span.End()

	span.SetAttributes(label.String("tweetId", tweetId))

	_, err := s.index.SaveObject(api.SearchModel{
		TweetID:  tweetId,
		FullText: fullText,
		Author:   author,
	})

	return err
}
