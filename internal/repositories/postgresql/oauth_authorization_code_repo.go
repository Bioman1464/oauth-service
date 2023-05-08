package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ ports.OauthAuthorizationCodeRepository = (*OauthAuthorizationCodeRepo)(nil)

func NewOauthAuthorizationCodeRepo(db *sql.DB) *OauthAuthorizationCodeRepo {
	return &OauthAuthorizationCodeRepo{db}
}

type OauthAuthorizationCodeRepo struct {
	db *sql.DB
}

func (o OauthAuthorizationCodeRepo) Create(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, code string, scopes string, expiresAt time.Time) error {
	//TODO implement me
	return nil
}

func (o OauthAuthorizationCodeRepo) GetByCode(ctx context.Context, code string) (domain.OauthAuthorizationCode, error) {
	//TODO implement me
	panic("implement me")
}
