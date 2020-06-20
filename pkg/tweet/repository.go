//go:generate mockgen -package=tweet -self_package=com.capturetweet/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"context"
	"gocloud.dev/docstore"
	"time"
)

type Repository interface {
	FindById(id string) (*Tweet, error)
	FindByIds(ids []string) ([]Tweet, error)
	Store(id, fullText, lang, userId string, retweetCount, favCount int, createdAt *time.Time) error
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

func (r repositoryImpl) Store(id, fullText, lang, userId string, retweetCount, favCount int, createdAt *time.Time) error {
	tweetCreatedAt := time.Now()
	if createdAt != nil {
		tweetCreatedAt = *createdAt
	}
	return r.coll.Create(context.Background(), &Tweet{
		ID:              id,
		CreatedAt:       tweetCreatedAt,
		UpdatedAt:       time.Now(),
		FullText:        fullText,
		CaptureURL:      nil, // TODO get url
		CaptureThumbURL: nil, // TODO get thumbnail ?
		Lang:            lang,
		FavoriteCount:   favCount,
		RetweetCount:    retweetCount,
		UserID:          userId,
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
