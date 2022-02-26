package user

import (
	"time"
)

type User struct {
	ID              string    `db:"id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdateAt        time.Time `db:"updated_at"`
	RegisterAt      time.Time `db:"registered_at"`
	Username        string    `db:"username"`
	ScreenName      string    `db:"screen_name"`
	Bio             string    `db:"bio"`
	ProfileImageURL string    `db:"profile_image_url"`
}
