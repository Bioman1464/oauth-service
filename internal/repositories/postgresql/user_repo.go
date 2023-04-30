package postgresql

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
