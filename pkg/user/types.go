package user

import (
	"time"
)

type User struct {
	ID              string    `docstore:"id"`
	CreatedAt       time.Time `docstore:"created_at"`
	UpdateAt        time.Time `docstore:"updated_at"`
	RegisterAt      time.Time `docstore:"registered_at"`
	Username        string    `docstore:"username"`
	ScreenName      string    `docstore:"screen_name"`
	Bio             string    `docstore:"bio"`
	ProfileImageURL string    `docstore:"profile_image_url"`

	DocstoreRevision interface{}
}
