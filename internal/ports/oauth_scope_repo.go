package ports

import (
	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthScopeRepository interface {
	GetByUUID(uuid uuid.UUID) (domain.OauthScope, error)
}
