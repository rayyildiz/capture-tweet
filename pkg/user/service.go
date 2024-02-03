package user

import (
	"context"
	"log/slog"
	"time"

	"capturetweet.com/api"
	"capturetweet.com/internal/convert"
	"github.com/ChimeraCoder/anaconda"
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
		slog.Error("user:service findById", slog.String("tweet_id", id), slog.Any("err", err))
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
		slog.Error("user:service findOrCreate", slog.String("tweet_id", id), slog.String("tweet_user", author.ScreenName), slog.Any("err", err))
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
