package user

import (
	"com.capturetweet/pkg/service"
	"github.com/jinzhu/gorm"
)

/*

type UserService interface {
	FindById(id string) (*UserModel, error)
	FindOrCreate(id, userName, screenName string) (*UserModel, error)
}
*/

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) service.UserService {
	return &serviceImpl{repo}
}

func (s serviceImpl) FindById(id string) (*service.UserModel, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return &service.UserModel{
		ID:         user.ID,
		UserName:   user.UserName,
		ScreenName: user.ScreenName,
	}, nil
}

func (s serviceImpl) FindOrCreate(id, userName, screenName string) (*service.UserModel, error) {
	user, err := s.repo.FindById(id)
	if user != nil {
		return &service.UserModel{
			ID:         user.ID,
			UserName:   user.UserName,
			ScreenName: user.ScreenName,
		}, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = s.repo.Store(id, userName, screenName)
	if err != nil {
		return nil, err
	}

	return &service.UserModel{
		ID:         id,
		UserName:   userName,
		ScreenName: screenName,
	}, nil
}
