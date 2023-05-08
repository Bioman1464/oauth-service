package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type OauthRefreshToken struct {
	ID                 uuid.UUID
	Token              string
	OauthAccessTokenID uuid.UUID
	ExpiresAt          time.Time
	CreatedAt          time.Time
	DeletedAt          null.Time
}
