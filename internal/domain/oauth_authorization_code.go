package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type OauthAuthorizationCode struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	OauthClientID uuid.UUID
	Code          string
	RedirectURL   string
	Scope         string
	ExpiresAt     time.Time
	CreatedAt     time.Time
	DeletedAt     null.Time
}
