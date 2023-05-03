package services

import (
	"context"

	"auth-service/internal/ports"
)

var _ OauthService = (*oauthService)(nil)

type OauthService interface {
	Authorize(ctx context.Context, userID string, clientID string, scopes []string) (string, error)
	CheckAuthorization(ctx context.Context) error
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

func (s *oauthService) Authorize(ctx context.Context, userID string, clientID string, scopes []string) (string, error) {
	//TODO::check session
	return "some code", nil
}

func (s *oauthService) CheckAuthorization(ctx context.Context) error {
	//TODO:: get user from session
	panic("Unimplemented")
}
