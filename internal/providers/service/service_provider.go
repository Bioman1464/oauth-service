package service

import (
	"auth-service/internal/providers/repo"
	"auth-service/internal/services"
)

type Provider interface {
	GetAuthService() *services.AuthService
	GetOauthService() *services.OauthService
}

type provider struct {
	auth  services.AuthService
	oauth services.OauthService
}

func NewServiceProvider(repoProvider repo.RepositoryProvider) Provider {
	return &provider{
		auth: services.NewAuthService(repoProvider.GetUserRepository()),
		oauth: services.NewOauthService(
			repoProvider.GetOauthAccessTokenRepository(),
			repoProvider.GetOauthClientRepository(),
			repoProvider.GetScopeRepository(),
			repoProvider.GetUserRepository(),
		),
	}
}

func (p *provider) GetAuthService() *services.AuthService {
	return &p.auth
}

func (p *provider) GetOauthService() *services.OauthService {
	return &p.oauth
}
