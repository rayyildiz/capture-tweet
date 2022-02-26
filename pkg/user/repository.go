//go:generate mockgen -package=user -self_package=capturetweet.com/pkg/user -destination=repository_mock.go . Repository
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByUserName(ctx context.Context, userName string) (*User, error)
	FindById(ctx context.Context, id string) (*User, error)

	Store(ctx context.Context, userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error
}

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) FindByUserName(ctx context.Context, userName string) (*User, error) {
	user := &User{}
	err := r.db.GetContext(ctx, user, `select * from users where username = $1`, userName)
	if err != nil {
		return nil, fmt.Errorf("error while getting data, %w", err)
	}
	return user, err
}

func (r repositoryImpl) FindById(ctx context.Context, id string) (*User, error) {
	user := &User{ID: id}
	err := r.db.GetContext(ctx, user, `select * from users where id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting data, %w", err)
	}
	return user, nil
}

func (r repositoryImpl) Store(ctx context.Context, userIdStr, userName, screenName, bio, profileImage string, registeredAt time.Time) error {

	_, err := r.db.ExecContext(ctx, "insert into users(id, username,screen_name,bio,profile_image_url, registered_at) values (?,?,?,?,?)", userIdStr, userName, screenName, bio, profileImage, registeredAt)

	if err != nil {
		return fmt.Errorf("wrror while inserting user, %w", err)
	}
	return nil
}
