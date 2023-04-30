package postgresql

import (
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

func (o OauthScopeRepository) GetByUUID(uuid uuid.UUID) (domain.OauthScope, error) {
	//TODO implement me
	panic("implement me")
}
