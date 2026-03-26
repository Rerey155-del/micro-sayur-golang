package seeds

import (
	"user-service/internal/core/domain/model"
	"user-service/utils/conv"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	// Memastikan Role "Customer" sudah ada
	modelRole := model.Role{}
	err := db.Where("name = ?", "Customer").First(&modelRole).Error
	if err != nil {
		log.Error().Err(err).Msg("Failed to get Customer role for seeding users")
		return
	}

	hashedPassword, _ := conv.HashPassword("user123")

	users := []model.User{
		{
			Name:       "Rerey Ganteng",
			Email:      "rerey@example.com",
			Password:   hashedPassword,
			Address:    "Jakarta, Indonesia",
			Phone:      "08123456789",
			IsVerified: true,
			Roles:      []model.Role{modelRole},
		},
		{
			Name:       "Siti Aminah",
			Email:      "siti@example.com",
			Password:   hashedPassword,
			Address:    "Bandung, West Java",
			Phone:      "08771234567",
			IsVerified: true,
			Roles:      []model.Role{modelRole},
		},
		{
			Name:       "Budi Santoso",
			Email:      "budi@example.com",
			Password:   hashedPassword,
			Address:    "Surabaya, East Java",
			Phone:      "08521234567",
			IsVerified: true,
			Roles:      []model.Role{modelRole},
		},
	}

	for _, u := range users {
		var count int64
		db.Model(&model.User{}).Where("email = ?", u.Email).Count(&count)
		if count == 0 {
			if err := db.Create(&u).Error; err != nil {
				log.Error().Err(err).Msgf("Failed to seed user %s", u.Name)
			}
		}
	}

	log.Info().Msg("Sample users seeded successfully")
}
