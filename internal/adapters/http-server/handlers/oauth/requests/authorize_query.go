package requests

import (
	"errors"

	"github.com/gin-gonic/gin"

	oauthErr "auth-service/internal/adapters/http-server/handlers/oauth/errors"
)

type AuthorizeQueryRequest struct {
	ResponseType string `json:"response_type" form:"response_type"`
	ClientID     string `json:"client_id" form:"client_id"`
	RedirectURL  string `json:"redirect_url" form:"redirect_url"`
	Scope        string `json:"scope" form:"scope"`
	State        string `json:"state" form:"state"`
}

func (r *AuthorizeQueryRequest) Parse(ctx *gin.Context) error {
	if err := ctx.BindQuery(r); err != nil {
		return err
	}

	return nil
}

func (r *AuthorizeQueryRequest) Validate() error {
	if r.RedirectURL == "" {
		return errors.New("redirect_url is not passed")
	}

	if r.ClientID == "" {
		return errors.New("client_id is not passed")
	}

	if r.ResponseType == "" {
		return errors.New("response type is not passed")
	}

	if r.ResponseType != "code" {
		return oauthErr.ErrUnsupportedResponseType
	}

	return nil
}
