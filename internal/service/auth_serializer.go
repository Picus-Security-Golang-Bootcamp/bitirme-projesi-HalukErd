package service

import (
	"BasketProjectGolang/internal/entity"
	"BasketProjectGolang/internal/model/api"
	"BasketProjectGolang/internal/util"
)

func signupRequestToEntity(request *api.SignupRequest) (*entity.User, error) {
	passwordBeforeHash := *request.Password
	hashedPassword, err := util.HashPassword(passwordBeforeHash)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Email:    *request.Email,
		Password: hashedPassword,
		Roles:    []string{"user"},
	}, nil
}

func createdUserToSignupResponse(user entity.User) *api.SignupResponse {
	return &api.SignupResponse{
		Id:    &user.ID,
		Email: &user.Email,
		Roles: user.Roles,
	}
}
