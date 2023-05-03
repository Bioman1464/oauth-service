package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"auth-service/internal/handlers/http-server/handlers"
	"auth-service/internal/handlers/http-server/handlers/auth/requests"
	"auth-service/internal/handlers/http-server/session"
	"auth-service/internal/services"
)

type Handler struct {
	handlers.BaseHandler
	authService services.AuthService
}

func NewAuthHandler(h handlers.BaseHandler, authService services.AuthService) *Handler {
	return &Handler{h, authService}
}

func (h *Handler) Login(ctx *gin.Context) {
	var req requests.LoginRequest

	if err := req.Parse(ctx); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "unable to parse request body",
		})

		return
	}

	user, err := h.authService.GetUser(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "user with passed credentials not found",
		})

		return
	}

	s := sessions.Default(ctx)
	s.Set(session.UserID, user.ID.String())

	err = s.Save()
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{
			"error":   err.Error(),
			"message": "unable to authorize user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// Register Not implemented
func (h *Handler) Register(ctx *gin.Context) {
	panic("Unimplemented")

	//TODO:: parse request
	//TODO:: validate request
	//TODO:: save user
	//TODO:: save user session
	//TODO:: response with success
}

func (h *Handler) GetUser(ctx *gin.Context) {
	s := sessions.Default(ctx)

	userID, ok := s.Get(session.UserID).(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user unauthorized"})

		return
	}

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{
			"error":   err.Error(),
			"message": "unable to identify user",
		})

		return
	}

	user, err := h.authService.GetUserByID(userUUID)
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{
			"error":   err.Error(),
			"message": "unable to identify user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) Logout(ctx *gin.Context) {
	s := sessions.Default(ctx)

	s.Delete(session.UserID)
	err := s.Save()
	if err != nil {
		h.Log.Error(err)

		ctx.JSON(http.StatusOK, gin.H{"message": "unable to logout"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
