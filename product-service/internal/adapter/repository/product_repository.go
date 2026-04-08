package repository

import (
	"context"
	"product-service/internal/core/domain/entity"
	"product-service/internal/core/domain/model"

	"gorm.io/gorm"
)

// Kontrak yang harus dipatuhi oleh repository
type ProductRepositoryInterface interface {
	GetAllProducts(ctx context.Context) ([]*entity.ProductEntity, error)
	CreateProduct(ctx context.Context, product *model.Product) error
}

// Struct yang mengimplementasikan kontrak di atas menggunakan GORM DB
type ProductRepository struct {
	DB *gorm.DB
}

// Fungsi untuk membuat objek repository baru
func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{DB: db}
}

// Logika untuk mengambil semua produk dari database
func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]*entity.ProductEntity, error) {
	var modelProducts []model.Product

	// Mengambil semua data produk menggunakan GORM
	err := r.DB.WithContext(ctx).Find(&modelProducts).Error
	if err != nil {
		return nil, err
	}

	// Konversi dari Model (Database) menjadi Entity (Aplikasi)
	var productEntities []*entity.ProductEntity
	for _, p := range modelProducts {
		productEntities = append(productEntities, &entity.ProductEntity{
			ID:          int64(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageURL:    p.ImageURL,
		})
	}

	return productEntities, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}
