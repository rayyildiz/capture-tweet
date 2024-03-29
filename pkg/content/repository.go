//go:generate go run go.uber.org/mock/mockgen@latest -package=content -self_package=capturetweet.com/pkg/content -destination=repository_mock.go . Repository
package content

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gocloud.dev/docstore"
)

type Repository interface {
	ContactUs(ctx context.Context, email, fullName, message string) error
}

type repositoryImpl struct {
	contactUs *docstore.Collection
}

func NewRepository(contactUs *docstore.Collection) Repository {
	return &repositoryImpl{
		contactUs: contactUs,
	}
}

func (r repositoryImpl) ContactUs(ctx context.Context, email, fullName, message string) error {
	id := uuid.New().String()
	return r.contactUs.Create(ctx, &ContactUs{
		ID:        id,
		CreatedAt: time.Now(),
		Email:     email,
		FullName:  fullName,
		Message:   message,
	})
}
