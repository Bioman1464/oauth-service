package ports

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type OauthAuthorizationCodeRepository interface {
	Create(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, code string, scopes string, expiresAt time.Time) error
	GetByCode(ctx context.Context, code string) (domain.OauthAuthorizationCode, error)
}
