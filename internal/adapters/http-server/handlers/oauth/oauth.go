package oauth

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"auth-service/internal/adapters/http-server/handlers"
	oauthErr "auth-service/internal/adapters/http-server/handlers/oauth/errors"
	"auth-service/internal/adapters/http-server/handlers/oauth/requests"
	"auth-service/internal/adapters/http-server/handlers/oauth/responses"
	"auth-service/internal/adapters/http-server/session"
	"auth-service/internal/services"
)

type Handler struct {
	handlers.BaseHandler
	oauthService services.OauthService
}

func NewOauthHandler(h handlers.BaseHandler, oauthService services.OauthService) *Handler {
	return &Handler{h, oauthService}
}

func (h *Handler) Authorize(ctx *gin.Context) {
	var req requests.AuthorizeQueryRequest

	if err := req.Parse(ctx); err != nil {
		ctx.Redirect(http.StatusFound, "/auth/error")

		return
	}

	if err := req.Validate(); err != nil {
		errResponse(ctx, oauthErr.ErrInvalidRequest, err)

		return
	}

	s := sessions.Default(ctx)
	userID, ok := s.Get(session.UserID).(string)
	if !ok {
		ctx.Redirect(http.StatusFound, "/auth/login")

		return
	}

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		errResponse(ctx, oauthErr.ErrServerError, err)

		return
	}

	clientUUID, err := uuid.FromString(req.ClientID)
	if err != nil {
		errResponse(ctx, oauthErr.ErrServerError, err)

		return
	}

	code, err := h.oauthService.Authorize(ctx, userUUID, clientUUID, req.Scope)
	if err != nil {
		errResponse(ctx, oauthErr.ErrServerError, err)

		return
	}

	response := responses.AuthorizationResponse{
		Code:  code,
		State: req.State,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) Token(ctx *gin.Context) {
	var req requests.TokenRequest

	if err := req.Parse(ctx); err != nil {
		errResponse(ctx, oauthErr.ErrInvalidRequest, err)

		return
	}

	if err := req.Validate(); err != nil {
		errResponse(ctx, oauthErr.ErrInvalidRequest, err)

		return
	}

	//TODO:: validate request
	//TODO:: get validate clientID, clientSecret, auth code
	//TODO:: generate access token, refresh token

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  "some token",
		"token_type":    "Bearer",
		"expires_id":    time.Now(),
		"refresh_token": "some refresh token",
		"scope":         "some scope",
	})
}

func (h *Handler) Validate(ctx *gin.Context) {
	//TODO:: get token
	//TODO:: validate token

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func errResponse(ctx *gin.Context, fallBackErr error, err error) {
	if oauthErr.IsHandledError(err) {
		errDefaultResponse(ctx, err)

		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":             fallBackErr.Error(),
		"error_description": err.Error(),
	})
}

func errDefaultResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":             err.Error(),
		"error_description": oauthErr.GetDescription(err),
	})
}
