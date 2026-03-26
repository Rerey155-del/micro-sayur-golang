package seeds

import (
	"product-service/internal/core/domain/model"
	"gorm.io/gorm"
)

func SeedProduct(db *gorm.DB) {
	products := []model.Product{
		{
			Name:        "Bayam Hijau",
			Description: "Bayam segar pilihan, kaya akan zat besi.",
			Price:       5000,
			Stock:       50,
		},
		{
			Name:        "Wortel Import",
			Description: "Wortel manis dan renyah, cocok untuk sop.",
			Price:       12000,
			Stock:       30,
		},
		{
			Name:        "Tomat Merah",
			Description: "Tomat matang pohon, segar dan berair.",
			Price:       8000,
			Stock:       100,
		},
	}

	for _, p := range products {
		// Cek jika produk sudah ada berdasarkan nama untuk menghindari duplikasi
		var count int64
		db.Model(&model.Product{}).Where("name = ?", p.Name).Count(&count)
		if count == 0 {
			db.Create(&p)
		}
	}
}
