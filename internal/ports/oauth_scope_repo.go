package ports

import (
	"context"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthScopeRepository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.OauthScope, error)
	Exists(ctx context.Context, scopes []string) (bool, error)
}
