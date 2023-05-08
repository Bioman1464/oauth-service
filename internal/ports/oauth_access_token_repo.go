package ports

import (
	"context"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthAccessTokenRepository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.OauthAccessToken, error)
}
