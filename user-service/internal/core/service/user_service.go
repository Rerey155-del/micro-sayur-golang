package service

import (
	"context"
	"errors"
	"user-service/internal/adapter/repository"
	"user-service/utils/conv"

	"github.com/labstack/gommon/log"
)

// UserServiceInterface mendefinisikan kontrak logika bisnis untuk user.
// Service bertanggung jawab untuk aturan bisnis, seperti validasi password, enkripsi, dll.
type UserServiceInterface interface {
	SignIn(ctx context.Context, email, password string) (string, error)
}

// userService adalah implementasi dari UserServiceInterface.
// Struct ini berinteraksi langsung dengan repository untuk akses ke database.
type userService struct {
	repo repository.UserRepositoryInterface
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

	// 3. Jika email ada dan password benar, kembalikan pesan sukses.
	// Nantinya di sini bisa ditambahkan logika untuk pembuatan token JWT.
	return "Login successful", nil
}

// NewUserService adalah fungsi generator untuk membuat instance service baru.
func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}
