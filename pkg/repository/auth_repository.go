package repository

import (
	"BasketProjectGolang/internal/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	repository := &authRepository{db: db}
	zap.L().Info("AuthRepository has been initialized.")
	//if err := repository.Migration(); err != nil {
	//	log.Fatalln("Auth Migration Failed")
	//}

	return repository
}

func (r *authRepository) Signup(user *entity.User) (*entity.User, error) {
	zap.L().Debug("repo.auth.create", zap.Reflect("user", user))
	if err := r.db.Create(user).Error; err != nil {
		zap.L().Error("repo.auth.create failed to Create user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (r *authRepository) GetUser(username string) (*entity.User, error) {
	return nil, nil
}

func (r *authRepository) Migration() error {
	return r.db.AutoMigrate(&entity.User{})
}
