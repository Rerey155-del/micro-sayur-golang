package seeds

import "gorm.io/gorm"

func SeedAdmin(db *gorm.DB) {
	bytes, err := conv.HashPasssword("admin123")
	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	admin := model.User{
		
	}
}
