package postgresql

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ ports.OauthAccessTokenRepository = (*OauthAccessTokenRepository)(nil)

func NewOauthAccessTokenRepository(db *sql.DB) *OauthAccessTokenRepository {
	return &OauthAccessTokenRepository{db}
}

type OauthAccessTokenRepository struct {
	db *sql.DB
}

func (o OauthAccessTokenRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.OauthAccessToken, error) {
	//TODO implement me
	panic("implement me")
}
