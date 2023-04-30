package requests

import (
	"github.com/gin-gonic/gin"
)

type AuthorizeQueryRequest struct {
	ClientId    string `json:"client_id"`
	RedirectURL string `json:"redirect_url"`
	Scopes      string `json:"scopes"`
	State       string `json:"state"`
}

func (r *AuthorizeQueryRequest) Parse(ctx *gin.Context) error {
	if err := ctx.BindQuery(r); err != nil {
		return err
	}

	return nil
}
