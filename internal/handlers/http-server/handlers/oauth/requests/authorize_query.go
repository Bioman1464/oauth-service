package requests

import (
	"github.com/gin-gonic/gin"

	"auth-service/internal/handlers/http-server/handlers/oauth/errors"
)

type AuthorizeQueryRequest struct {
	ClientID    string `json:"client_id" form:"client_id"`
	RedirectURL string `json:"redirect_url" form:"redirect_url"`
	Scopes      string `json:"scopes" form:"scopes"`
	State       string `json:"state" form:"state"`
}

func (r *AuthorizeQueryRequest) Parse(ctx *gin.Context) error {
	if err := ctx.BindQuery(r); err != nil {
		return err
	}

	return nil
}

func (r *AuthorizeQueryRequest) Validate() error {
	if r.RedirectURL == "" || r.Scopes == "" || r.ClientID == "" {
		return errors.ErrInvalidRequest
	}

	return nil
}
