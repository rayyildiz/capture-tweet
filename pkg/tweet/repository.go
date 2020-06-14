package tweet

import (
	"com.capturetweet/pkg/model"
	"github.com/jinzhu/gorm"
	"time"
)

type Repository interface {
	FindById(id string) (*model.Tweet, error)
	Store(id, fullText, lang string, retweetCount, favCount int, createdAt *time.Time) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) FindById(id string) (*model.Tweet, error) {
	return nil, nil
}

func (r repositoryImpl) Store(id, fullText, lang string, retweetCount, favCount int, createdAt *time.Time) error {

	return nil
}
