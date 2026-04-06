package service

import (
	"context"
	"errors"
	"user-service/config"
	"user-service/internal/adapter/repository"
	"user-service/internal/core/domain/entity"
	"user-service/utils/conv"
	"user-service/utils/jwt"

	"github.com/labstack/gommon/log"
)

// UserServiceInterface mendefinisikan kontrak logika bisnis untuk user.
// Service bertanggung jawab untuk aturan bisnis, seperti validasi password, enkripsi, dll.
type UserServiceInterface interface {
	SignIn(ctx context.Context, email, password string) (string, error)
	FetchAllUsers(ctx context.Context) ([]entity.UserEntity, error)
}

// userService adalah implementasi dari UserServiceInterface.
// Struct ini berinteraksi langsung dengan repository untuk akses ke database.
type userService struct {
	repo repository.UserRepositoryInterface
	cfg  *config.Config
}

// SignIn menjalankan logika untuk proses login user.
func (u *userService) SignIn(ctx context.Context, email, password string) (string, error) {
	// 1. Meminta repository untuk mencari data user berdasarkan email.
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		// Log error jika user tidak ditemukan atau terjadi kesalahan database.
		log.Errorf("[UserService-1] SignIn: %v", err)
		return "", err
	}

	// 2. Membandingkan password yang diinput user dengan password (hash) di database.
	// Kita tidak menyimpan password dalam bentuk teks biasa (plain text), tapi dalam bentuk hash (enkripsi).
	if checkPass := conv.CheckPasswordHash(password, user.Password); !checkPass {
		err = errors.New("password is incorrect")
		log.Errorf("[UserService-1] SignIn: %v", err)
		return "", err
	}

	// 3. Jika email ada dan password benar, generate token JWT.
	token, err := jwt.GenerateToken(*user, u.cfg.Jwt.JwtSecretKey, u.cfg.Jwt.JwtIssuer)
	if err != nil {
		log.Errorf("[UserService-3] SignIn: %v", err)
		return "", err
	}

	return token, nil
}

// FetchAllUsers mengambil semua data user.
func (u *userService) FetchAllUsers(ctx context.Context) ([]entity.UserEntity, error) {
	users, err := u.repo.FetchAllUsers(ctx)
	if err != nil {
		log.Errorf("[UserService] FetchAllUsers: %v", err)
		return nil, err
	}
	return users, nil
}

// NewUserService adalah fungsi generator untuk membuat instance service baru.
func NewUserService(repo repository.UserRepositoryInterface, cfg *config.Config) UserServiceInterface {
	return &userService{
		repo: repo,
		cfg:  cfg,
	}
}
