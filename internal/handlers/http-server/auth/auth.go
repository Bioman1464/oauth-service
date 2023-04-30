package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/handlers/http-server/auth/requests"
	"auth-service/internal/services"
)

type Handler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *Handler {
	return &Handler{authService}
}

func (h *Handler) Login(ctx *gin.Context) {
	var req requests.LoginRequest

	if err := req.Parse(ctx); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

	//TODO:: save user session
}

func (h *Handler) Register(ctx *gin.Context) {
	var req requests.LoginRequest

	if err := req.Parse(ctx); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})

		return
	}

	//TODO:: save user session
	//TODO:: response with success
}
