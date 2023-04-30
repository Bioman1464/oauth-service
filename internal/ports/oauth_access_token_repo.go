package ports

import (
	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthAccessTokenRepository interface {
	GetByUUID(uuid uuid.UUID) (domain.OauthAccessToken, error)
}
