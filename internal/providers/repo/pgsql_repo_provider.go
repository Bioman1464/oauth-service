package repo

import (
	"database/sql"

	"auth-service/internal/ports"
	"auth-service/internal/repositories/postgresql"
)

type RepositoryProvider interface {
	GetOauthAccessTokenRepository() ports.OauthAccessTokenRepository
	GetOauthClientRepository() ports.OauthClientRepository
	GetOauthScopeRepository() ports.OauthScopeRepository
	GetUserRepository() ports.UserRepository
	GetOauthAuthorizationCodeRepository() ports.OauthAuthorizationCodeRepository
}

var _ RepositoryProvider = (*repositoryProvider)(nil)

type repositoryProvider struct {
	oauthAccessToken       ports.OauthAccessTokenRepository
	oauthClient            ports.OauthClientRepository
	oauthScope             ports.OauthScopeRepository
	user                   ports.UserRepository
	oauthAuthorizationCode ports.OauthAuthorizationCodeRepository
}

func NewPostgresqlRepositoryProvider(db *sql.DB) RepositoryProvider {
	return &repositoryProvider{
		oauthAccessToken:       postgresql.NewOauthAccessTokenRepository(db),
		oauthClient:            postgresql.NewOauthClientRepository(db),
		oauthScope:             postgresql.NewOauthScopeRepository(db),
		user:                   postgresql.NewUserRepository(db),
		oauthAuthorizationCode: postgresql.NewOauthAuthorizationCodeRepo(db),
	}
}

func (p *repositoryProvider) GetOauthAccessTokenRepository() ports.OauthAccessTokenRepository {
	return p.oauthAccessToken
}

func (p *repositoryProvider) GetOauthClientRepository() ports.OauthClientRepository {
	return p.oauthClient
}

func (p *repositoryProvider) GetOauthScopeRepository() ports.OauthScopeRepository {
	return p.oauthScope
}

func (p *repositoryProvider) GetUserRepository() ports.UserRepository {
	return p.user
}

func (p *repositoryProvider) GetOauthAuthorizationCodeRepository() ports.OauthAuthorizationCodeRepository {
	return p.oauthAuthorizationCode
}
