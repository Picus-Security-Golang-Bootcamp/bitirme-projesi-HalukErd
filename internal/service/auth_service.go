package service

import (
	"BasketProjectGolang/internal/model/api"
	"BasketProjectGolang/internal/repository"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthService interface {
	Signup(ctx context.Context, req *api.SignupRequest) (*api.SignupResponse, error)
	PermissionCheck(g *gin.Context, permission string) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *authService {
	a := &authService{repo: repo}
	zap.L().Info("AuthService has been initialized.")
	return a
}

func (a authService) Signup(ctx context.Context, req *api.SignupRequest) (*api.SignupResponse, error) {
	user, err := signupRequestToEntity(req)
	if err != nil {
		return nil, err
	}

	creteadUser, err := a.repo.Signup(user)
	signupResponse := createdUserToSignupResponse(*creteadUser)
	return signupResponse, nil
}

func (a authService) PermissionCheck(g *gin.Context, permission string) error {
	ctxRolesValue, _ := g.Get("roles")
	roles := ctxRolesValue.([]string)
	for _, role := range roles {
		if role == permission {
			return nil
		}
	}
	return errors.New("there is no permission for granted roles")
}
