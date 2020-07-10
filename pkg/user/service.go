package user

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/convert"
	"github.com/ChimeraCoder/anaconda"
	"go.uber.org/zap"
	"time"
)

type serviceImpl struct {
	repo Repository
	log  *zap.Logger
}

func NewService(repo Repository, log *zap.Logger) api.UserService {
	return &serviceImpl{repo, log}
}

func (s serviceImpl) FindById(id string) (*api.UserModel, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		s.log.Error("user:service findById", zap.String("tweet_id", id), zap.Error(err))
		return nil, err
	}

	return &api.UserModel{
		ID:           user.ID,
		UserName:     user.Username,
		ScreenName:   user.ScreenName,
		Bio:          user.Bio,
		ProfileImage: convert.String(user.ProfileImageURL),
	}, nil
}

func (s serviceImpl) FindOrCreate(author *anaconda.User) (*api.UserModel, error) {
	id := author.IdStr

	user, err := s.repo.FindById(id)
	if user != nil {
		return &api.UserModel{
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

	err = s.repo.Store(id, author.ScreenName, author.Name, author.Description, author.ProfileImageUrlHttps, registeredAt)
	if err != nil {
		s.log.Error("user:service findOrCreate", zap.String("tweet_id", id), zap.String("tweet_user", author.ScreenName), zap.Error(err))
		return nil, err
	}

	return &api.UserModel{
		ID:           id,
		UserName:     author.ScreenName,
		ScreenName:   author.Name,
		Bio:          author.Description,
		ProfileImage: convert.String(author.ProfileImageURL),
	}, nil
}
