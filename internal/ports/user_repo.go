package ports

import (
	"context"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
)

type UserRepository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.User, error)
	Create(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	DeleteByUUID(ctx context.Context, uuid uuid.UUID) error
}
