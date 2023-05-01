package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/handlers/http-server/oauth/requests"
	"auth-service/internal/services"
)

type Handler struct {
	oauthService services.OauthService
}

func NewOauthHandler(oauthService services.OauthService) *Handler {
	return &Handler{oauthService}
}

func (h *Handler) Authorize(ctx *gin.Context) {
	var req requests.AuthorizeQueryRequest

	if err := req.Parse(ctx); err != nil {
		//TODO:: redirect to front error page

		return
	}

	if req.RedirectURL == "" {
		ctx.Header("Cache-Control", "no-store")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})

		return
	}

	if req.Scopes == "" || req.ClientId == "" {
		ctx.Header("Cache-Control", "no-store")
		ctx.Redirect(http.StatusFound, req.RedirectURL)
		return
	}

	//TODO:: validate scopes, clientID
	//TODO:: get userID from session
	//TODO:: check authorization if none -> login
	//TODO::

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) Token(ctx *gin.Context) {
	//TODO:: validate request
	//TODO:: get validate clientID, clientSecret, auth code
	//TODO:: generate access token, refresh token

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) Validate(ctx *gin.Context) {
	//TODO:: get token
	//TODO:: validate token

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
