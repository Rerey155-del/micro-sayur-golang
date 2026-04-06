package seeds

import (
	"user-service/internal/core/domain/model"
	"user-service/utils/conv"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	SeedRole(db)

	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to hash password")
	}

	modelRole := model.Role{}
	err = db.Where("name = ?", "Super Admin").First(&modelRole).Error
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get role")
	}

	admin := model.User{
		Name:       "Super Admin",
		Email:      "superadmin@gmail.com",
		Password:   bytes,
		Roles:      []model.Role{modelRole},
		IsVerified: true,
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "superadmin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to seed admin user")
	}

	// Tambahan Admin Baru untuk Anda gunakan login
	hashedPassword, _ := conv.HashPassword("password123")
	newAdmin := model.User{
		Name:       "Rerey Admin",
		Email:      "rerey@admin.com",
		Password:   hashedPassword,
		Roles:      []model.Role{modelRole},
		IsVerified: true,
	}

	if err := db.FirstOrCreate(&newAdmin, model.User{Email: "rerey@admin.com"}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to seed additional admin user")
	}

	log.Info().Msg("Admin users seeded successfully")
}
