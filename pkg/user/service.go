package user

import (
	"context"
	"time"

	"capturetweet.com/api"
	"capturetweet.com/internal/convert"
	"github.com/ChimeraCoder/anaconda"
	"go.uber.org/zap"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) api.UserService {
	return &serviceImpl{
		repo: repo,
	}
}

func (s serviceImpl) FindById(ctx context.Context, id string) (*api.UserModel, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		zap.L().Error("user:service findById", zap.String("tweet_id", id), zap.Error(err))
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

func (s serviceImpl) FindOrCreate(ctx context.Context, author *anaconda.User) (*api.UserModel, error) {
	id := author.IdStr

	user, err := s.repo.FindById(ctx, id)
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

	err = s.repo.Store(ctx, id, author.ScreenName, author.Name, author.Description, author.ProfileImageUrlHttps, registeredAt)
	if err != nil {
		zap.L().Error("user:service findOrCreate", zap.String("tweet_id", id), zap.String("tweet_user", author.ScreenName), zap.Error(err))
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
