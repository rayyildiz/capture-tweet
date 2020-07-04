//go:generate mockgen -package=tweet -self_package=com.capturetweet/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"context"
	"github.com/ChimeraCoder/anaconda"
	"gocloud.dev/docstore"
	"gocloud.dev/gcerrors"
	"io"
	"time"
)

type Repository interface {
	FindById(id string) (*Tweet, error)
	FindByIds(ids []string) ([]Tweet, error)
	Store(tweet *anaconda.Tweet) error
	Exist(id string) bool
	UpdateLargeImage(id, captureUrl string) error
	UpdateThumbImage(id, captureUrl string) error

	FindAllOrderByUpdated(size int) ([]Tweet, error)
}

type repositoryImpl struct {
	coll *docstore.Collection
}

func NewRepository(coll *docstore.Collection) Repository {
	return &repositoryImpl{coll}
}

func (r repositoryImpl) FindById(id string) (*Tweet, error) {
	tweet := &Tweet{ID: id}
	err := r.coll.Get(context.Background(), tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (r repositoryImpl) Exist(id string) bool {
	tweet := &Tweet{ID: id}

	err := r.coll.Get(context.Background(), tweet, "id")
	code := gcerrors.Code(err)
	if code == gcerrors.NotFound {
		return false
	}
	return true
}

func (r repositoryImpl) Store(tweet *anaconda.Tweet) error {
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

	return r.coll.Create(context.Background(), &Tweet{
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

func (r repositoryImpl) FindByIds(ids []string) ([]Tweet, error) {
	var list []Tweet

	for _, id := range ids {
		if tweet, err := r.FindById(id); err == nil {
			list = append(list, *tweet)
		}
	}
	return list, nil
}

func (r repositoryImpl) UpdateLargeImage(id, captureUrl string) error {
	tweet := &Tweet{ID: id}
	return r.coll.Actions().Update(tweet, docstore.Mods{"capture_url": captureUrl, "updated_at": time.Now()}).Do(context.Background())
}

func (r repositoryImpl) UpdateThumbImage(id, captureUrl string) error {
	tweet := &Tweet{ID: id}
	return r.coll.Actions().Update(tweet, docstore.Mods{"capture_thumb_url": captureUrl, "updated_at": time.Now()}).Do(context.Background())
}

func (r repositoryImpl) FindAllOrderByUpdated(size int) ([]Tweet, error) {
	var tweets []Tweet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
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
