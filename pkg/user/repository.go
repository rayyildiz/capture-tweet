//go:generate mockgen -package=user -self_package=capturetweet.com/pkg/user -destination=repository_mock.go . Repository
package user

import (
	"capturetweet.com/internal/ent"
	"capturetweet.com/internal/ent/user"
	"context"
	"fmt"
	"time"
)

type Repository interface {
	FindByUserName(ctx context.Context, userName string) (*User, error)
	FindById(ctx context.Context, id string) (*User, error)

	Store(ctx context.Context, userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error
}

type repositoryImpl struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) Repository {
	return &repositoryImpl{db}
}

func toUser(u *ent.User) *User {
	return &User{
		ID:              u.ID,
		CreatedAt:       u.CreatedAt,
		UpdateAt:        u.UpdatedAt,
		RegisterAt:      u.RegisteredAt,
		Username:        u.Username,
		ScreenName:      u.ScreenName,
		Bio:             u.Bio,
		ProfileImageURL: u.ProfileImageURL,
	}
}

func (r repositoryImpl) FindByUserName(ctx context.Context, userName string) (*User, error) {
	u, err := r.db.User.Query().Where(user.Username(userName)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting data, %w", err)
	}
	return toUser(u), nil
}

func (r repositoryImpl) FindById(ctx context.Context, id string) (*User, error) {
	u, err := r.db.User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting data, %w", err)
	}
	return toUser(u), nil
}

func (r repositoryImpl) Store(ctx context.Context, userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error {
	_, err := r.db.User.Create().SetID(userIdStr).SetUsername(userName).SetScreenName(screenName).SetBio(bio).SetProfileImageURL(profileImage).SetRegisteredAt(registeredAt).Save(ctx)

	if err != nil {
		return fmt.Errorf("wrror while inserting user, %w", err)
	}
	return nil
}
