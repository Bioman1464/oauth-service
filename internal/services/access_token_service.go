package services

import (
	"context"

	"auth-service/internal/ports"
)

var _ OauthService = (*oauthService)(nil)

type OauthService interface {
	Authorize(ctx context.Context)
	CheckAuthorization(ctx context.Context)
}

type oauthService struct {
	accessTokenRepo ports.OauthAccessTokenRepository
	clientRepo      ports.OauthClientRepository
	scopeRepo       ports.OauthScopeRepository
	userRepo        ports.UserRepository
}

func NewOauthService(
	accessTokenRepo ports.OauthAccessTokenRepository,
	clientRepo ports.OauthClientRepository,
	scopeRepo ports.OauthScopeRepository,
	userRepo ports.UserRepository,
) OauthService {
	return &oauthService{
		accessTokenRepo: accessTokenRepo,
		clientRepo:      clientRepo,
		scopeRepo:       scopeRepo,
		userRepo:        userRepo,
	}
}

func (s *oauthService) Authorize(ctx context.Context) {
	//TODO::check session
}

func (s *oauthService) CheckAuthorization(ctx context.Context) {
	//TODO:: get user from session
}
