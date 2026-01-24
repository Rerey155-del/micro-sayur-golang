package seeds

import (
	"log"
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
			log.Fatalf("%s:%v", err.Error(), err)
		} else {
			log.Printf("Role %s created", role.Name)
		}
	}
}

func SeedAdmin(db *gorm.DB) {
	SeedRole(db)
	// Tambahkan seed data lainnya di sini jika perlu
}