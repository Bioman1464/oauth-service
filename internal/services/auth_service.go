package services

import (
	"github.com/gofrs/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/ports"
)

var _ AuthService = (*authService)(nil)

type AuthService interface {
	GetUser(email string, password string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
}

type authService struct {
	userRepo ports.UserRepository
}

func NewAuthService(
	userRepo ports.UserRepository,
) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) GetUser(email string, password string) (*domain.User, error) {
	id, _ := uuid.NewV7()

	return &domain.User{
		ID:        id,
		FirstName: "test",
		LastName:  "test",
		Email:     "email@mail.com",
	}, nil
}

func (a *authService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	return &domain.User{
		ID:        id,
		FirstName: "test",
		LastName:  "test",
		Email:     "email@email.com",
	}, nil
}
