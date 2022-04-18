package api

import "github.com/gin-gonic/gin"

type SignupRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (req SignupRequest) validate(g *gin.Context) error {
	return nil
}
