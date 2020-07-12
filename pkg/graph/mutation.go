package graph

import (
	"com.capturetweet/internal/convert"
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

type mutationResolverImpl struct {
}

func newMutationResolver() MutationResolver {
	return &mutationResolverImpl{}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url string) (*Tweet, error) {
	id, err := _twitterService.Store(ctx, url)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("capture error", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	model, err := _twitterService.FindById(ctx, id)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("capture error, findById", zap.String("id", id), zap.Error(err))
		return nil, err
	}
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
	err := _contentService.SendMail(ctx, input.Email, input.FullName, input.Message)
	if err != nil {
		sentry.CaptureException(err)
		_log.Error("could not send mail", zap.String("contact_email", input.Email), zap.String("contact_fullName", input.FullName), zap.Error(err))
		return "", errors.New("error occurred, please try again or contact from info@capturetweet.com mail address")
	}
	return "we saved your message and we will contact you as soon as possible, thanks for feedback", nil
}
