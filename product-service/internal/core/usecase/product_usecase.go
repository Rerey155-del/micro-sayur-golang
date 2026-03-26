package usecase

import (
	"context"
	"product-service/internal/adapter/repository"
	"product-service/internal/core/domain/entity"
)

// Kontrak untuk logika bisnis produk
type ProductUsecaseInterface interface {
	FetchAllProducts(ctx context.Context) ([]*entity.ProductEntity, error)
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
