package handler

import (
	"order-service/internal/core/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	usecase usecase.OrderUsecaseInterface
}

type createOrderRequest struct {
	UserID       int64                    `json:"user_id"`
	CustomerName string                   `json:"customer_name"`
	Address      string                   `json:"address"`
	Notes        string                   `json:"notes"`
	Items        []createOrderItemRequest `json:"items"`
}

type createOrderItemRequest struct {
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

func NewOrderHandler(router fiber.Router, uc usecase.OrderUsecaseInterface) {
	handler := &OrderHandler{usecase: uc}

	api := router.Group("/api/v1")
	api.Get("/orders", handler.GetOrders)
	api.Get("/orders/user/:userId", handler.GetOrdersByUserID)
	api.Post("/orders", handler.CreateOrder)
	api.Patch("/orders/:id/status", handler.UpdateOrderStatus)
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orders, err := h.usecase.FetchAllOrders(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    orders,
	})
}

func (h *OrderHandler) GetOrdersByUserID(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("userId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	orders, err := h.usecase.FetchOrdersByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    orders,
	})
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	req := createOrderRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	items := make([]usecase.CreateOrderItemInput, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, usecase.CreateOrderItemInput{
			ProductID:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     item.Quantity,
		})
	}

	err := h.usecase.CreateOrder(c.Context(), usecase.CreateOrderInput{
		UserID:       req.UserID,
		CustomerName: req.CustomerName,
		Address:      req.Address,
		Notes:        req.Notes,
		Items:        items,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
	})
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	req := updateStatusRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.usecase.UpdateOrderStatus(c.Context(), uint(id), req.Status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order status updated successfully",
	})
}
