//go:generate mockgen -package=tweet -self_package=capturetweet.com/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"capturetweet.com/internal/convert"
	"capturetweet.com/internal/ent"
	"capturetweet.com/internal/ent/schema"
	"capturetweet.com/internal/ent/tweet"
	"context"
	"fmt"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type Repository interface {
	FindById(ctx context.Context, id string) (*Tweet, error)
	FindByIds(ctx context.Context, ids []string) ([]Tweet, error)
	FindByUser(ctx context.Context, userId string) ([]Tweet, error)
	Store(ctx context.Context, tweet *anaconda.Tweet) error
	Exist(ctx context.Context, id string) bool
	UpdateLargeImage(ctx context.Context, id, captureUrl string) error
	UpdateThumbImage(ctx context.Context, id, captureUrl string) error

	FindAllOrderByUpdated(ctx context.Context, size int) ([]Tweet, error)
}

type repositoryImpl struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) Repository {
	return &repositoryImpl{db}
}

// FindById returns a tweet by id. Return err if not found.
func (r repositoryImpl) FindById(ctx context.Context, id string) (*Tweet, error) {
	t, err := r.db.Tweet.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting from database %w", err)
	}

	return &Tweet{
		ID:              t.ID,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		PostedAt:        t.PostedAt,
		FullText:        t.FullText,
		CaptureURL:      convert.String(t.CaptureURL),
		CaptureThumbURL: convert.String(t.CaptureThumbURL),
		Lang:            t.Lang,
		FavoriteCount:   t.FavoriteCount,
		RetweetCount:    t.RetweetCount,
		AuthorID:        t.AuthorID,
		Resources:       toResource(t),
	}, nil
}

func toResource(t *ent.Tweet) []Resource {
	var list []Resource

	for _, r := range t.Resources {
		list = append(list, Resource{
			URL:       r.URL,
			Width:     r.Width,
			Height:    r.Height,
			MediaType: r.MediaType,
		})
	}
	return list
}

func toTweet(l *ent.Tweet) Tweet {
	return Tweet{
		ID:              l.ID,
		CreatedAt:       l.CreatedAt,
		UpdatedAt:       l.UpdatedAt,
		PostedAt:        l.PostedAt,
		FullText:        l.FullText,
		CaptureURL:      convert.String(l.CaptureURL),
		CaptureThumbURL: convert.String(l.CaptureThumbURL),
		Lang:            l.Lang,
		FavoriteCount:   l.FavoriteCount,
		RetweetCount:    l.RetweetCount,
		AuthorID:        l.AuthorID,
		Resources:       toResource(l),
	}
}

// Exist returns true if the tweet exists. Otherwise, return false.
func (r repositoryImpl) Exist(ctx context.Context, id string) bool {
	exist, err := r.db.Tweet.Query().Where(tweet.ID(id)).Exist(ctx)

	if err != nil {
		return false
	}

	return exist
}

// Store a tweet.
func (r repositoryImpl) Store(ctx context.Context, tweet *anaconda.Tweet) error {
	postedAt, err := tweet.CreatedAtTime()
	if err != nil {
		postedAt = time.Now()
	}
	var resources []schema.Resource
	for _, media := range tweet.Entities.Media {
		resources = append(resources, schema.Resource{
			URL:       media.Media_url_https,
			Width:     media.Sizes.Medium.H,
			Height:    media.Sizes.Medium.W,
			MediaType: media.Type,
		})
	}
	_, err = r.db.Tweet.Create().SetID(tweet.IdStr).
		SetPostedAt(postedAt).
		SetFullText(tweet.FullText).
		SetLang(tweet.Lang).
		SetResources(resources).
		SetCreatedAt(time.Now()).
		SetCaptureURL("").
		SetCaptureThumbURL("").
		SetFavoriteCount(tweet.FavoriteCount).
		SetRetweetCount(tweet.RetweetCount).
		SetAuthorID(tweet.User.IdStr).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("error while inserting tweet %w", err)
	}

	return nil
}

// FindByIds returns a list of tweets by ids. Return err if not found.
func (r repositoryImpl) FindByIds(ctx context.Context, ids []string) ([]Tweet, error) {
	var tweets []Tweet

	list, err := r.db.Tweet.Query().Where(tweet.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error wh,le getting tweets, %w", err)
	}

	for _, l := range list {
		tweets = append(tweets, toTweet(l))
	}

	return tweets, nil
}

// FindByUser returns a list of tweets by user id.
func (r repositoryImpl) FindByUser(ctx context.Context, userId string) ([]Tweet, error) {
	var tweets []Tweet

	list, err := r.db.Tweet.Query().Where(tweet.AuthorID(userId)).Limit(24).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting tweets, %w", err)
	}

	for _, l := range list {
		tweets = append(tweets, toTweet(l))
	}
	return tweets, nil
}

// UpdateLargeImage updates the large image of the tweet in the database.
func (r repositoryImpl) UpdateLargeImage(ctx context.Context, id, captureUrl string) error {
	_, err := r.db.Tweet.Update().SetCaptureURL(captureUrl).Where(tweet.ID(id)).Save(ctx)

	return err
}

// UpdateThumbImage updates the thumb image of the tweet in the database.
func (r repositoryImpl) UpdateThumbImage(ctx context.Context, id, captureUrl string) error {
	_, err := r.db.Tweet.Update().SetCaptureThumbURL(captureUrl).Where(tweet.ID(id)).Save(ctx)

	return err
}

// FindAllOrderByUpdated returns a list of tweets ordered by updated_at.
func (r repositoryImpl) FindAllOrderByUpdated(ctx context.Context, size int) ([]Tweet, error) {
	var tweets []Tweet
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	list, err := r.db.Tweet.Query().Order(ent.Desc(tweet.FieldPostedAt)).Limit(24).All(ctx)
	if err != nil {
		return nil, err
	}

	for _, l := range list {
		tweets = append(tweets, toTweet(l))
	}

	return tweets, nil
}
