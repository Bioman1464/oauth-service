package repo

import (
	"database/sql"

	"auth-service/internal/ports"
	"auth-service/internal/repositories/postgresql"
)

type RepositoryProvider interface {
	GetOauthAccessTokenRepository() ports.OauthAccessTokenRepository
	GetOauthClientRepository() ports.OauthClientRepository
	GetScopeRepository() ports.OauthScopeRepository
	GetUserRepository() ports.UserRepository
}

var _ RepositoryProvider = (*repositoryProvider)(nil)

type repositoryProvider struct {
	oauthAccessToken ports.OauthAccessTokenRepository
	oauthClient      ports.OauthClientRepository
	oauthScope       ports.OauthScopeRepository
	user             ports.UserRepository
}

func NewPostgresqlRepositoryProvider(db *sql.DB) RepositoryProvider {
	return &repositoryProvider{
		oauthAccessToken: postgresql.NewOauthAccessTokenRepository(db),
		oauthClient:      postgresql.NewOauthClientRepository(db),
		oauthScope:       postgresql.NewOauthScopeRepository(db),
		user:             postgresql.NewUserRepository(db),
	}
}

func (p *repositoryProvider) GetOauthAccessTokenRepository() ports.OauthAccessTokenRepository {
	return p.oauthAccessToken
}

func (p *repositoryProvider) GetOauthClientRepository() ports.OauthClientRepository {
	return p.oauthClient
}

func (p *repositoryProvider) GetScopeRepository() ports.OauthScopeRepository {
	return p.oauthScope
}

func (p *repositoryProvider) GetUserRepository() ports.UserRepository {
	return p.user
}
