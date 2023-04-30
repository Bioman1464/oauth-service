package postgresql

import (
	"database/sql"

	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ ports.OauthClientRepository = (*OauthClientRepository)(nil)

func NewOauthClientRepository(db *sql.DB) *OauthClientRepository {
	return &OauthClientRepository{db}
}

type OauthClientRepository struct {
	db *sql.DB
}

func (o OauthClientRepository) GetByUUID(uuid uuid.UUID) (domain.OauthClient, error) {
	//TODO implement me
	panic("implement me")
}
