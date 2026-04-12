package repository

import (
	"context"
	"order-service/internal/core/domain/entity"
	"order-service/internal/core/domain/model"

	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	GetAllOrders(ctx context.Context) ([]*entity.OrderEntity, error)
	GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.OrderEntity, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrderStatus(ctx context.Context, id uint, status string) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepositoryInterface {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]*entity.OrderEntity, error) {
	var orders []model.Order
	if err := r.DB.WithContext(ctx).Preload("Items").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}

	return mapOrdersToEntities(orders), nil
}

func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.OrderEntity, error) {
	var orders []model.Order
	if err := r.DB.WithContext(ctx).
		Preload("Items").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return mapOrdersToEntities(orders), nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	return r.DB.WithContext(ctx).Create(order).Error
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, id uint, status string) error {
	return r.DB.WithContext(ctx).Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func mapOrdersToEntities(orders []model.Order) []*entity.OrderEntity {
	result := make([]*entity.OrderEntity, 0, len(orders))

	for _, order := range orders {
		items := make([]*entity.OrderItemEntity, 0, len(order.Items))
		for _, item := range order.Items {
			items = append(items, &entity.OrderItemEntity{
				ID:           int64(item.ID),
				OrderID:      int64(item.OrderID),
				ProductID:    item.ProductID,
				ProductName:  item.ProductName,
				ProductPrice: item.ProductPrice,
				Quantity:     item.Quantity,
				Subtotal:     item.Subtotal,
			})
		}

		result = append(result, &entity.OrderEntity{
			ID:           int64(order.ID),
			OrderCode:    order.OrderCode,
			UserID:       order.UserID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			TotalAmount:  order.TotalAmount,
			Address:      order.Address,
			Notes:        order.Notes,
			Items:        items,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}

	return result
}
