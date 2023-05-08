package services

import (
	"context"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	oauthErr "auth-service/internal/adapters/http-server/handlers/oauth/errors"
	"auth-service/internal/ports"
	"auth-service/pkg/random"
)

var _ OauthService = (*oauthService)(nil)

type OauthService interface {
	Authorize(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, scope string) (string, error)
	CheckAuthorization(ctx context.Context) error
}

type oauthService struct {
	accessTokenRepo       ports.OauthAccessTokenRepository
	clientRepo            ports.OauthClientRepository
	scopeRepo             ports.OauthScopeRepository
	userRepo              ports.UserRepository
	authorizationCodeRepo ports.OauthAuthorizationCodeRepository
}

func NewOauthService(
	accessTokenRepo ports.OauthAccessTokenRepository,
	clientRepo ports.OauthClientRepository,
	scopeRepo ports.OauthScopeRepository,
	userRepo ports.UserRepository,
	authorizationCodeRepo ports.OauthAuthorizationCodeRepository,
) OauthService {
	return &oauthService{
		accessTokenRepo:       accessTokenRepo,
		clientRepo:            clientRepo,
		scopeRepo:             scopeRepo,
		userRepo:              userRepo,
		authorizationCodeRepo: authorizationCodeRepo,
	}
}

func (s *oauthService) Authorize(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, scope string) (string, error) {
	//Validate client
	exists, err := s.clientRepo.Exists(ctx, clientID)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", oauthErr.ErrUnauthorizedClient
	}

	//Validate scope
	if scope != "" {
		scopes := strings.Split(scope, " ")

		exists, err = s.scopeRepo.Exists(ctx, scopes)
		if err != nil {
			return "", err
		}

		if !exists {
			return "", oauthErr.ErrInvalidScope
		}
	}

	code := random.GenerateRandomString(16)
	expiresAt := time.Now().Add(time.Minute * 10)
	if err = s.authorizationCodeRepo.Create(ctx, userID, clientID, code, scope, expiresAt); err != nil {
		return "", err
	}

	return code, nil
}

func (s *oauthService) CheckAuthorization(ctx context.Context) error {
	//TODO:: get user from session
	panic("Unimplemented")
}

func (s *oauthService) GetToken(ctx context.Context, clientID string) error {
	return nil
}
