package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
