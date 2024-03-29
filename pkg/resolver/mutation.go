package resolver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"capturetweet.com/internal/convert"
)

type mutationResolverImpl struct {
}

func newMutationResolver() MutationResolver {
	return &mutationResolverImpl{}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url string) (*Tweet, error) {
	id, err := _twitterService.Store(ctx, url)
	if err != nil {
		slog.Error("capture error", slog.String("url", url), slog.Any("err", err))
		return nil, err
	}

	model, err := _twitterService.FindById(ctx, id)
	if err != nil {
		slog.Error("capture error, findById", slog.String("id", id), slog.Any("err", err))
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

func (r mutationResolverImpl) Contact(ctx context.Context, input ContactInput, tweetID *string, captcha string) (string, error) {
	msg := input.Message
	if tweetID != nil {
		msg = fmt.Sprintf("Tweet ID: %s\n Message: %s", *tweetID, input.Message)
	}

	err := _contentService.StoreContactRequest(ctx, input.Email, input.FullName, msg, captcha)
	if err != nil {
		slog.Error("could not send mail", slog.String("contact_email", input.Email), slog.String("contact_fullName", input.FullName), slog.Any("err", err))
		return "", errors.New("error occurred, please try again or contact from info@capturetweet.com mail address")
	}
	return "We saved your message and we will contact you as soon as possible, thanks for feedback", nil
}
