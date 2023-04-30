package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Parse(ctx *gin.Context) error {
	if err := ctx.BindJSON(r); err != nil {
		return err
	}

	return nil
}

func (r *LoginRequest) Validate() error {
	if r.Password == "" || r.Email == "" {
		return errors.New("required fields are not filled")
	}

	return nil
}
