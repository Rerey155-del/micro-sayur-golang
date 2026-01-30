package repository

import (
	"context"

	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var modelUser model.User

	err := r.DB.WithContext(ctx).
		Where("email = ? AND is_verified = ?", email, true).
		Preload("Roles").
		First(&modelUser).Error

	if err != nil {
		log.Error().Err(err).Msg("[UserRepository] GetUserByEmail failed")
		return nil, err
	}

	return &entity.UserEntity{
		ID:         int64(modelUser.ID),
		Name:       modelUser.Name,
		Email:      modelUser.Email,
		Password:   modelUser.Password,
		RoleName:   modelUser.Roles[0].Name,
		Address:    modelUser.Address,
		Phone:      modelUser.Phone,
		Photo:      modelUser.Photo,
		Lat:        modelUser.Lat,
		Lng:        modelUser.Lng,
		IsVerified: modelUser.IsVerified,
	}, nil

}
