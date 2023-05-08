package ports

import (
	"context"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthClientRepository interface {
	Exists(ctx context.Context, uuid uuid.UUID) (bool, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.OauthClient, error)
}
