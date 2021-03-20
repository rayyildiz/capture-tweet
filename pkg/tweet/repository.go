//go:generate mockgen -package=tweet -self_package=com.capturetweet/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"context"
	"github.com/ChimeraCoder/anaconda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/docstore"
	"gocloud.dev/gcerrors"
	"io"
	"sort"
	"time"
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
	coll   *docstore.Collection
	tracer trace.Tracer
}

func NewRepository(coll *docstore.Collection) Repository {
	return &repositoryImpl{
		coll:   coll,
		tracer: otel.GetTracerProvider().Tracer("com.capturetweet/pkg/tweet"),
	}
}

func (r repositoryImpl) FindById(ctx context.Context, id string) (*Tweet, error) {
	ctx, span := r.tracer.Start(ctx, "repo-findById")
	defer span.End()

	span.AddEvent("findById", trace.WithAttributes(attribute.String("tweetId", id)))

	tweet := &Tweet{ID: id}
	err := r.coll.Get(ctx, tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (r repositoryImpl) Exist(ctx context.Context, id string) bool {
	ctx, span := r.tracer.Start(ctx, "repo-exist")
	defer span.End()

	span.AddEvent("exist", trace.WithAttributes(attribute.String("tweetId", id)))

	tweet := &Tweet{ID: id}

	err := r.coll.Get(ctx, tweet, "id")
	code := gcerrors.Code(err)
	if code == gcerrors.NotFound {
		return false
	}
	return true
}

func (r repositoryImpl) Store(ctx context.Context, tweet *anaconda.Tweet) error {
	ctx, span := r.tracer.Start(ctx, "repo-store")
	defer span.End()

	span.AddEvent("store", trace.WithAttributes(attribute.String("tweetId", tweet.IdStr)))

	postedAt, err := tweet.CreatedAtTime()
	if err != nil {
		postedAt = time.Now()
	}
	var resources []Resource
	for _, media := range tweet.Entities.Media {
		resources = append(resources, Resource{
			ID:        media.Id_str,
			URL:       media.Media_url_https,
			Width:     media.Sizes.Medium.H,
			Height:    media.Sizes.Medium.W,
			MediaType: media.Type,
		})
	}
	return r.coll.Create(ctx, &Tweet{
		ID:              tweet.IdStr,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		PostedAt:        postedAt,
		FullText:        tweet.FullText,
		CaptureURL:      nil,
		CaptureThumbURL: nil,
		Lang:            tweet.Lang,
		FavoriteCount:   tweet.FavoriteCount,
		RetweetCount:    tweet.RetweetCount,
		AuthorID:        tweet.User.IdStr,
		Resources:       resources,
	})
}

func (r repositoryImpl) FindByIds(ctx context.Context, ids []string) ([]Tweet, error) {
	ctx, span := r.tracer.Start(ctx, "repo-findByIds")
	defer span.End()

	span.AddEvent("findByIds")

	var list []Tweet

	for _, id := range ids {
		if tweet, err := r.FindById(ctx, id); err == nil {
			list = append(list, *tweet)
		}
	}

	sort.Sort(SortByPosted(list))

	return list, nil
}
func (r repositoryImpl) FindByUser(ctx context.Context, userId string) ([]Tweet, error) {
	ctx, span := r.tracer.Start(ctx, "repo-findByUser")
	defer span.End()
	span.AddEvent("findByUser", trace.WithAttributes(attribute.String("userId", userId)))

	iterator := r.coll.Query().Where("author_id", "=", userId).Limit(24).Get(ctx)
	defer iterator.Stop()

	var tweets []Tweet
	for {
		var tweet Tweet
		err := iterator.Next(ctx, &tweet)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		tweets = append(tweets, tweet)
	}

	sort.Sort(SortByPosted(tweets))

	return tweets, nil
}

func (r repositoryImpl) UpdateLargeImage(ctx context.Context, id, captureUrl string) error {
	ctx, span := r.tracer.Start(ctx, "repo-updateLargeImage")
	defer span.End()

	span.AddEvent("updateLargeImage", trace.WithAttributes(attribute.String("tweetId", id)))

	tweet := &Tweet{ID: id}
	return r.coll.Actions().Update(tweet, docstore.Mods{"capture_url": captureUrl, "updated_at": time.Now()}).Do(ctx)
}

func (r repositoryImpl) UpdateThumbImage(ctx context.Context, id, captureUrl string) error {
	ctx, span := r.tracer.Start(ctx, "repo-updateThumbImage")
	defer span.End()
	span.AddEvent("updateThumbImage", trace.WithAttributes(attribute.String("tweetId", id)))

	tweet := &Tweet{ID: id}
	return r.coll.Actions().Update(tweet, docstore.Mods{"capture_thumb_url": captureUrl, "updated_at": time.Now()}).Do(ctx)
}

func (r repositoryImpl) FindAllOrderByUpdated(ctx context.Context, size int) ([]Tweet, error) {
	ctx, span := r.tracer.Start(ctx, "repo-findAllOrderByUpdated")
	defer span.End()
	span.AddEvent("findAllOrderByUpdated")

	var tweets []Tweet
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	it := r.coll.Query().Limit(size).OrderBy("updated_at", "desc").Get(ctx)

	for {
		var tweet Tweet
		err := it.Next(ctx, &tweet)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
