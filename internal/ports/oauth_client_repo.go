package ports

import (
	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthClientRepository interface {
	GetByUUID(uuid uuid.UUID) (domain.OauthClient, error)
}
