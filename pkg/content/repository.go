//go:generate mockgen -package=content -self_package=capturetweet.com/pkg/content -destination=repository_mock.go . Repository
package content

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	ContactUs(ctx context.Context, email, fullName, message string) error
}

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) ContactUs(ctx context.Context, email, fullName, message string) error {
	id := uuid.New().String()
	_, err := r.db.ExecContext(ctx, `insert into contact_us(id, email, full_name, message, created_at)  values($1,$2,$3,$4,$5)`, id, email, fullName, message, time.Now())

	return err
}
