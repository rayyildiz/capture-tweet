package content

import (
	"time"
)

type ContactUs struct {
	ID               string    `docstore:"id"`
	CreatedAt        time.Time `docstore:"created_at"`
	Email            string    `docstore:"email"`
	FullName         string    `docstore:"full_name"`
	Message          string    `docstore:"message"`
	DocstoreRevision interface{}
}
