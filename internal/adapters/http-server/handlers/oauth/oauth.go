package oauth

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"auth-service/internal/adapters/http-server/handlers"
	"auth-service/internal/adapters/http-server/handlers/oauth/requests"
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
		//TODO:: redirect to front error page

		return
	}

	if err := req.Validate(); err != nil {
		h.Log.Info(req)

		ctx.Header("Cache-Control", "no-store")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})

		return
	}

	s := sessions.Default(ctx)
	userID, ok := s.Get(session.UserID).(string)
	if !ok {
		ctx.Header("Cache-Control", "no-store")
		ctx.Redirect(http.StatusFound, "/auth/login")

		return
	}

	scopeArr := strings.Split(req.State, ",")

	code, err := h.oauthService.Authorize(ctx, userID, req.ClientID, scopeArr)
	if err != nil {
		ctx.Header("Cache-Control", "no-store")
		ctx.JSON(http.StatusBadRequest, "some error")

		return
	}

	//TODO:: validate scopes, clientID
	//TODO:: get userID from session

	data := make(map[string]string)
	data["code"] = code
	data["state"] = req.State

	ctx.JSON(http.StatusOK, gin.H{"data": data})
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
