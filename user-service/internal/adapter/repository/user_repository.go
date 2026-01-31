package repository

import (
	"context"
	"errors"

	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"gorm.io/gorm"
)

// UserRepositoryInterface mendefinisikan kontrak untuk akses data user ke database.
// Layer repository fokus hanya pada operasi database (Query, Insert, Update, Delete).
type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

// UserRepository adalah implementasi dari UserRepositoryInterface menggunakan GORM.
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository adalah fungsi generator untuk membuat instance repository baru.
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

// GetUserByEmail mencari data user di database berdasarkan alamat email.
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	// modelUser digunakan untuk menampung hasil query database (mapping ke tabel database).
	var modelUser model.User

	// r.DB.WithContext(ctx) memastikan query menghormati timeout atau pembatalan dari context.
	// .Where() menambahkan filter pencarian.
	// .Preload("Roles") otomatis mengambil data role user dari tabel user_roles (join).
	// .First(&modelUser) mengambil satu record pertama yang ditemukan.
	err := r.DB.WithContext(ctx).
		Where("email = ? AND is_verified = ?", email, true).
		Preload("Roles").
		First(&modelUser).Error

	if err != nil {
		// Jika data tidak ditemukan, GORM mengembalikan error gorm.ErrRecordNotFound.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Mengonversi (mapping) dari model database ke entity domain.
	// Ini dilakukan untuk memisahkan struktur database dengan struktur data aplikasi internal.
	return &entity.UserEntity{
		ID:         int64(modelUser.ID),
		Name:       modelUser.Name,
		Email:      modelUser.Email,
		Password:   modelUser.Password,
		RoleName:   modelUser.Roles[0].Name, // Mengambil nama role pertama dari hasil preload.
		Address:    modelUser.Address,
		Phone:      modelUser.Phone,
		Photo:      modelUser.Photo,
		Lat:        modelUser.Lat,
		Lng:        modelUser.Lng,
		IsVerified: modelUser.IsVerified,
	}, nil

}
