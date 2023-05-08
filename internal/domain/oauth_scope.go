package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type OauthScope struct {
	id        uuid.UUID
	value     string
	createdAt time.Time
}
