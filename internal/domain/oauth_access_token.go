package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type OauthAccessToken struct {
	ID                         uuid.UUID
	Token                      string
	CurrentOauthRefreshTokenID uuid.UUID
	Scope                      string
	CreatedAt                  time.Time
	ExpiresAt                  time.Time
	DeletedAt                  null.Time
}
