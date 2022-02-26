package content

import (
	"time"
)

type ContactUs struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Email     string    `db:"email"`
	FullName  string    `db:"full_name"`
	Message   string    `db:"message"`
}
