//go:generate mockgen -package=tweet -self_package=com.capturetweet/pkg/tweet -destination=repository_mock.go . Repository
package tweet

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/model"
	"github.com/jinzhu/gorm"
	"time"
)

type Repository interface {
	FindById(id string) (*model.Tweet, error)
	FindByIds(ids []string) ([]model.Tweet, error)
	Store(id, fullText, lang, userId string, retweetCount, favCount int, createdAt *time.Time) error

	FindBySearch(term string, limit int, cursorId *string) ([]model.Tweet, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) FindById(id string) (*model.Tweet, error) {
	tweet := &model.Tweet{}

	err := r.db.Where(&model.Tweet{ID: id}).First(tweet).Error
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
	return r.db.Create(&model.Tweet{
		ID:              id,
		CreatedAt:       tweetCreatedAt,
		UpdatedAt:       time.Now(),
		FullText:        fullText,
		CaptureURL:      "", // TODO get url
		CaptureThumbURL: "", // TODO get thumbnail ?
		Lang:            lang,
		FavoriteCount:   favCount,
		RetweetCount:    retweetCount,
		UserID:          userId,
	}).Error
}

func (r repositoryImpl) FindBySearch(term string, limit int, cursorId *string) ([]model.Tweet, error) {
	if cursorId == nil {
		cursorId = convert.String("")
	}
	var list []model.Tweet

	err := r.db.Where("id > & AND full_text is '%?%'", cursorId, term).Limit(limit).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r repositoryImpl) FindByIds(ids []string) ([]model.Tweet, error) {

}
