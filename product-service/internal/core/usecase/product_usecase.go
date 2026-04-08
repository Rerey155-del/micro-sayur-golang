package usecase

import (
	"context"
	"product-service/internal/adapter/repository"
	"product-service/internal/core/domain/entity"
	productModel "product-service/internal/core/domain/model"
)

// Kontrak untuk logika bisnis produk
type ProductUsecaseInterface interface {
	FetchAllProducts(ctx context.Context) ([]*entity.ProductEntity, error)
	CreateProduct(ctx context.Context, name, description string, price float64, stock int, imageURL string) error
}

type ProductUsecase struct {
	repo repository.ProductRepositoryInterface
}

func NewProductUsecase(repo repository.ProductRepositoryInterface) ProductUsecaseInterface {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) FetchAllProducts(ctx context.Context) ([]*entity.ProductEntity, error) {
	// Di sini Anda bisa menambahkan logika bisnis tambahan (misal: caching, filter, dll)
	return uc.repo.GetAllProducts(ctx)
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, name, description string, price float64, stock int, imageURL string) error {
	productModel := &productModel.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		ImageURL:    imageURL,
	}
	return uc.repo.CreateProduct(ctx, productModel)
}
