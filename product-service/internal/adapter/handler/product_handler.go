package handler

import (
	"product-service/internal/core/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase usecase.ProductUsecaseInterface
}

func NewProductHandler(router fiber.Router, uc usecase.ProductUsecaseInterface) {
	handler := &ProductHandler{usecase: uc}

	api := router.Group("/api/v1")
	api.Get("/products", handler.GetProducts)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.usecase.FetchAllProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    products,
	})
}
