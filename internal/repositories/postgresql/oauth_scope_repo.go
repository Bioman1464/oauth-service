package postgresql

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ ports.OauthScopeRepository = (*OauthScopeRepository)(nil)

func NewOauthScopeRepository(db *sql.DB) *OauthScopeRepository {
	return &OauthScopeRepository{db}
}

type OauthScopeRepository struct {
	db *sql.DB
}

func (o OauthScopeRepository) Exists(ctx context.Context, scopes []string) (bool, error) {
	//TODO implement me
	return true, nil
}

func (o OauthScopeRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.OauthScope, error) {
	//TODO implement me
	panic("implement me")
}
