package seeds

import (
	"user-service/internal/core/domain/model"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []model.Role{
		{
			Name: "Super Admin",
		},
		{
			Name: "Customer",
		},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			log.Fatal().Err(err).Msgf("Failed to seed role %s", role.Name)
		} else {
			log.Info().Msgf("Role %s created", role.Name)
		}
	}
}
