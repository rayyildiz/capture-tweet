package user

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/service"
	"github.com/ChimeraCoder/anaconda"
	"time"
)

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
		ID:           user.ID,
		UserName:     user.Username,
		ScreenName:   user.ScreenName,
		Bio:          user.Bio,
		ProfileImage: convert.String(user.ProfileImageURL),
	}, nil
}

func (s serviceImpl) FindOrCreate(author *anaconda.User) (*service.UserModel, error) {
	id := author.IdStr

	user, err := s.repo.FindById(id)
	if user != nil {
		return &service.UserModel{
			ID:           user.ID,
			UserName:     user.Username,
			ScreenName:   user.ScreenName,
			Bio:          author.Description,
			ProfileImage: convert.String(author.ProfileImageURL),
		}, nil
	}

	registeredAt, err := time.Parse(time.RubyDate, author.CreatedAt)
	if err != nil {
		registeredAt = time.Now()
	}

	err = s.repo.Store(id, author.ScreenName, author.Name, author.Description, author.ProfileBackgroundImageUrlHttps, registeredAt)
	if err != nil {
		return nil, err
	}

	return &service.UserModel{
		ID:           id,
		UserName:     author.ScreenName,
		ScreenName:   author.Name,
		Bio:          author.Description,
		ProfileImage: convert.String(author.ProfileImageURL),
	}, nil
}
