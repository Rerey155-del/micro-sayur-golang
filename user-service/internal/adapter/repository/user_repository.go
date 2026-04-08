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
	FetchAllUsers(ctx context.Context) ([]entity.UserEntity, error)
	CreateUser(ctx context.Context, user *entity.UserEntity) error
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

// FetchAllUsers mengambil semua data user dari database.
func (r *UserRepository) FetchAllUsers(ctx context.Context) ([]entity.UserEntity, error) {
	var modelUsers []model.User
	err := r.DB.WithContext(ctx).Preload("Roles").Find(&modelUsers).Error
	if err != nil {
		return nil, err
	}

	var userEntities []entity.UserEntity
	for _, m := range modelUsers {
		roleName := ""
		if len(m.Roles) > 0 {
			roleName = m.Roles[0].Name
		}

		userEntities = append(userEntities, entity.UserEntity{
			ID:         int64(m.ID),
			Name:       m.Name,
			Email:      m.Email,
			RoleName:   roleName,
			Address:    m.Address,
			Phone:      m.Phone,
			IsVerified: m.IsVerified,
		})
	}

	return userEntities, nil
}

// CreateUser mendaftarkan user baru ke dalam database dan mengaitkannya dengan Role.
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.UserEntity) error {
	// Dapatkan role berdasarkan RoleName
	var role model.Role
	err := r.DB.WithContext(ctx).Where("name = ?", user.RoleName).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("role not found")
		}
		return err
	}

	modelUser := model.User{
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: true, // Auto verifikasi untuk kemudahan
		Roles:      []model.Role{role},
	}

	return r.DB.WithContext(ctx).Create(&modelUser).Error
}
