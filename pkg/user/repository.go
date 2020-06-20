//go:generate mockgen -package=user -self_package=com.capturetweet/pkg/user -destination=repository_mock.go . Repository
package user

import (
	"context"
	"gocloud.dev/docstore"
	"time"
)

type Repository interface {
	FindByUserName(userName string) (*User, error)
	FindById(id string) (*User, error)

	Store(userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error
}

type repositoryImpl struct {
	coll *docstore.Collection
}

func NewRepository(coll *docstore.Collection) Repository {
	return &repositoryImpl{coll}
}

func (r repositoryImpl) FindByUserName(userName string) (*User, error) {

	user := &User{}
	it := r.coll.Query().Where("username", "=", userName).Limit(1).Get(context.Background())

	err := it.Next(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r repositoryImpl) FindById(id string) (*User, error) {
	user := &User{ID: id}
	err := r.coll.Get(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r repositoryImpl) Store(userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error {
	return r.coll.Put(context.Background(), &User{
		ID:              userIdStr,
		CreatedAt:       time.Now(),
		UpdateAt:        time.Now(),
		RegisterAt:      registeredAt,
		Username:        userName,
		ScreenName:      screenName,
		Bio:             bio,
		ProfileImageURL: profileImage,
	})
}
