package requests

import "github.com/gin-gonic/gin"

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (r *RegisterRequest) Parse(ctx *gin.Context) error {
	if err := ctx.BindJSON(r); err != nil {
		return err
	}

	return nil
}
