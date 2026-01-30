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
	} else {
		log.Info().Msg("Admin user seeded successfully")
	}
}
