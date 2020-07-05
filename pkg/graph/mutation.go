package graph

import (
	"com.capturetweet/internal/convert"
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.uber.org/zap"
)

type mutationResolverImpl struct {
}

func newMutationResolver() MutationResolver {
	return &mutationResolverImpl{}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url string) (*Tweet, error) {
	tr := global.Tracer("capturetweet/api")
	spanCtx, span := tr.Start(ctx, "capture")
	defer span.End()

	id, err := _twitterService.Store(url)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("capture error", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	span.AddEvent(spanCtx, "tweet captured", kv.String("id", id))

	model, err := _twitterService.FindById(id)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("capture error, findById", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	span.AddEvent(spanCtx, "tweet captured, get details", kv.String("id", id))

	var resources []*Resource
	for _, res := range model.Resources {
		resources = append(resources, &Resource{
			ID:        res.ID,
			URL:       res.URL,
			MediaType: convert.String(res.ResourceType),
			Width:     convert.Int(res.Width),
			Height:    convert.Int(res.Height),
		})
	}

	return &Tweet{
		ID:              model.ID,
		FullText:        model.FullText,
		PostedAt:        model.PostedAt,
		AuthorID:        convert.String(model.AuthorID),
		CaptureURL:      model.CaptureURL,
		CaptureThumbURL: model.CaptureThumbURL,
		FavoriteCount:   convert.Int(model.FavoriteCount),
		Lang:            convert.String(model.Lang),
		RetweetCount:    convert.Int(model.RetweetCount),
		Resources:       resources,
	}, nil
}

func (r mutationResolverImpl) Contact(ctx context.Context, input ContactInput) (string, error) {
	tr := global.Tracer("capturetweet/api")
	spanCtx, span := tr.Start(ctx, "capture")
	defer span.End()

	err := _contentService.SendMail(input.Email, input.FullName, input.Message)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("could not send mail", zap.String("contact_email", input.Email), zap.String("contact_fullName", input.FullName), zap.Error(err))
		span.AddEvent(spanCtx, "could not insert", kv.String("email", input.Email))
		return "", errors.New("error occurred, please try again or contact from info@capturetweet.com mail address")
	}
	span.AddEvent(spanCtx, "successfully stored the contact", kv.String("email", input.Email))
	return "we saved your message and we will contact you as soon as possible, thanks for feedback", nil
}
