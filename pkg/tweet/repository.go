//go:generate mockgen -package=tweet -self_package=capturetweet.com/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
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
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db}
}

// FindById returns a tweet by id. Return err if not found.
func (r repositoryImpl) FindById(ctx context.Context, id string) (*Tweet, error) {
	tweet := &Tweet{}
	err := r.db.GetContext(ctx, tweet, `select * from tweets where id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting from database %w", err)
	}

	return tweet, nil
}

// Exist returns true if the tweet exists. Otherwise, return false.
func (r repositoryImpl) Exist(ctx context.Context, id string) bool {
	var count int

	err := r.db.GetContext(ctx, &count, `select count(*) from tweets where id=$1`, id)
	if err != nil {
		return false
	}

	return count == 0
}

// Store a tweet.
func (r repositoryImpl) Store(ctx context.Context, tweet *anaconda.Tweet) error {
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

	t := &Tweet{
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
	}
	_, err = r.db.NamedExecContext(ctx, `insert into tweets(id, full_text, capture_url, capture_thumb_url, lang, author_id, posted_at) values(:id, :full_text, :capture_url, :capture_thumb_url, :lang, :author_id, :posted_at) `, t)
	return err
}

// FindByIds returns a list of tweets by ids. Return err if not found.
func (r repositoryImpl) FindByIds(ctx context.Context, ids []string) ([]Tweet, error) {
	var tweets []Tweet

	err := r.db.SelectContext(ctx, &tweets, `select * from tweets where id in(?) order by posted_at`, ids)
	if err != nil {
		return nil, fmt.Errorf("error while getting tweets, %w", err)
	}

	return tweets, nil
}

// FindByUser returns a list of tweets by user id.
func (r repositoryImpl) FindByUser(ctx context.Context, userId string) ([]Tweet, error) {

	var tweets []Tweet

	err := r.db.SelectContext(ctx, &tweets, `select * from tweets where author_id=$1 order by posted_at limit 24`, userId)
	if err != nil {
		return nil, fmt.Errorf("error while getting tweets, %w", err)
	}

	return tweets, nil
}

// UpdateLargeImage updates the large image of the tweet in the database.
func (r repositoryImpl) UpdateLargeImage(ctx context.Context, id, captureUrl string) error {
	_, err := r.db.ExecContext(ctx, `update tweets set capture_url=$2 AND updated_at=$3 where id=$1`, id, captureUrl, time.Now())

	return err
}

// UpdateThumbImage updates the thumb image of the tweet in the database.
func (r repositoryImpl) UpdateThumbImage(ctx context.Context, id, captureUrl string) error {
	_, err := r.db.ExecContext(ctx, `update tweets set capture_thumb_url=$2 AND updated_at=$3 where id=$1`, id, captureUrl, time.Now())

	return err
}

// FindAllOrderByUpdated returns a list of tweets ordered by updated_at.
func (r repositoryImpl) FindAllOrderByUpdated(ctx context.Context, size int) ([]Tweet, error) {
	var tweets []Tweet
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := r.db.SelectContext(ctx, tweets, `select * from tweets order by updated_at desc limit $1`, size)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
