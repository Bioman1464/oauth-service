package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (r *TokenRequest) Parse(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(r); err != nil {
		return err
	}

	return nil
}

func (r *TokenRequest) Validate() error {
	if r.ClientID == "" || r.ClientSecret == "" {
		return errors.New("client - id, secret credentials are not passed")
	}

	if r.RedirectURI == "" {
		return errors.New("redirect_uri is not passed")
	}

	if r.Code == "" {
		return errors.New("code is not passed")
	}

	if r.GrantType != "authorization_code" {
		return errors.New("grant type is not supported")
	}

	return nil
}
