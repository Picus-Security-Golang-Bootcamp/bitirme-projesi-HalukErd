package repository

import (
	"BasketProjectGolang/internal/entity"
)

type AuthRepository interface {
	Signup(user *entity.User) (*entity.User, error)
	GetUser(username string) (*entity.User, error)
}
