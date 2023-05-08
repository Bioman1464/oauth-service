package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type OauthClient struct {
	ID        uuid.UUID
	Label     string
	DomainURL string
	Secret    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
