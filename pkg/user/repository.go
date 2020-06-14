//go:generate mockgen -package=user -self_package=com.capturetweet/pkg/user -destination=repository_mock.go . Repository
package user

import (
	"com.capturetweet/pkg/model"
	"github.com/jinzhu/gorm"
	"time"
)

type Repository interface {
	FindByUserName(userName string) (*model.User, error)
	FindById(id string) (*model.User, error)

	Store(userIdStr, userName, screenName string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) FindByUserName(userName string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where(&model.User{UserName: userName}).First(user).Error
	return user, err
}

func (r repositoryImpl) FindById(id string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where(&model.User{ID: id}).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r repositoryImpl) Store(userIdStr, userName, screenName string) error {
	return r.db.Create(&model.User{
		ID:         userIdStr,
		UserName:   userName,
		ScreenName: screenName,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}).Error
}
