//go:generate mockgen -package=content -self_package=capturetweet.com/pkg/content -destination=repository_mock.go . Repository
package content

import (
	"capturetweet.com/internal/ent"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	ContactUs(ctx context.Context, email, fullName, message string) error
}

type repositoryImpl struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) ContactUs(ctx context.Context, email, fullName, message string) error {
	id := uuid.New().String()

	_, err := r.db.ContactUs.Create().SetID(id).SetMessage(message).SetFullName(fullName).SetEmail(email).Save(ctx)
	return err
}
