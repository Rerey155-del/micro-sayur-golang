package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"order-service/internal/adapter/repository"
	"order-service/internal/core/domain/entity"
	orderModel "order-service/internal/core/domain/model"
	"strings"
	"time"
)

type CreateOrderItemInput struct {
	ProductID    int64
	ProductName  string
	ProductPrice float64
	Quantity     int
}

type CreateOrderInput struct {
	UserID       int64
	CustomerName string
	Address      string
	Notes        string
	Items        []CreateOrderItemInput
}

type OrderUsecaseInterface interface {
	FetchAllOrders(ctx context.Context) ([]*entity.OrderEntity, error)
	FetchOrdersByUserID(ctx context.Context, userID int64) ([]*entity.OrderEntity, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) error
	UpdateOrderStatus(ctx context.Context, id uint, status string) error
}

type OrderUsecase struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderUsecase(repo repository.OrderRepositoryInterface) OrderUsecaseInterface {
	return &OrderUsecase{repo: repo}
}

func (uc *OrderUsecase) FetchAllOrders(ctx context.Context) ([]*entity.OrderEntity, error) {
	return uc.repo.GetAllOrders(ctx)
}

func (uc *OrderUsecase) FetchOrdersByUserID(ctx context.Context, userID int64) ([]*entity.OrderEntity, error) {
	return uc.repo.GetOrdersByUserID(ctx, userID)
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, input CreateOrderInput) error {
	if input.UserID <= 0 {
		return errors.New("user_id is required")
	}
	if strings.TrimSpace(input.CustomerName) == "" {
		return errors.New("customer_name is required")
	}
	if len(input.Items) == 0 {
		return errors.New("items are required")
	}

	orderItems := make([]orderModel.OrderItem, 0, len(input.Items))
	totalAmount := 0.0

	for _, item := range input.Items {
		if item.ProductID <= 0 || strings.TrimSpace(item.ProductName) == "" {
			return errors.New("invalid product item")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than zero")
		}
		if item.ProductPrice < 0 {
			return errors.New("product_price is invalid")
		}

		subtotal := item.ProductPrice * float64(item.Quantity)
		totalAmount += subtotal

		orderItems = append(orderItems, orderModel.OrderItem{
			ProductID:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     item.Quantity,
			Subtotal:     subtotal,
		})
	}

	order := &orderModel.Order{
		OrderCode:    generateOrderCode(),
		UserID:       input.UserID,
		CustomerName: strings.TrimSpace(input.CustomerName),
		Status:       "Pending",
		TotalAmount:  totalAmount,
		Address:      strings.TrimSpace(input.Address),
		Notes:        strings.TrimSpace(input.Notes),
		Items:        orderItems,
	}

	return uc.repo.CreateOrder(ctx, order)
}

func (uc *OrderUsecase) UpdateOrderStatus(ctx context.Context, id uint, status string) error {
	normalizedStatus := normalizeStatus(status)
	if normalizedStatus == "" {
		return errors.New("status is invalid")
	}

	return uc.repo.UpdateOrderStatus(ctx, id, normalizedStatus)
}

func generateOrderCode() string {
	now := time.Now()
	randomSuffix := rand.New(rand.NewSource(now.UnixNano())).Intn(9000) + 1000
	return fmt.Sprintf("ORD-%s-%d", now.Format("20060102"), randomSuffix)
}

func normalizeStatus(status string) string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "pending":
		return "Pending"
	case "diproses", "processing":
		return "Diproses"
	case "dikirim", "shipping":
		return "Dikirim"
	case "selesai", "success", "completed":
		return "Selesai"
	case "dibatalkan", "canceled", "cancelled":
		return "Dibatalkan"
	default:
		return ""
	}
}
