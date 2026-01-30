package service

import (
	"context"
	"errors"
	"user-service/internal/adapter/repository"
	"user-service/utils/conv"
	"github.com/labstack/gommon/log"
)

type UserServiceInterface interface {
	SignIn(ctx context.Context, email, password string) (string, error)
}

type userService struct {
	repo repository.UserRepositoryInterface
}

func (u *userService) SignIn(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Errorf("[UserService-1] SignIn: %v", err)
		return "", err
	}

	if checkPass := conv.CheckPasswordHash(password, user.Password); !checkPass {
		err = errors.New("password is incorrect")
		log.Errorf("[UserService-1] SignIn: %v", err)
		return "", err
	}

	return "Login successful", nil
}

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}
