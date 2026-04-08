package handler

import (
	"fmt"
	"path/filepath"
	"product-service/internal/core/usecase"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase usecase.ProductUsecaseInterface
}

func NewProductHandler(router fiber.Router, uc usecase.ProductUsecaseInterface) {
	handler := &ProductHandler{usecase: uc}

	api := router.Group("/api/v1")
	api.Get("/products", handler.GetProducts)
	api.Post("/products", handler.CreateProduct)
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

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")

	// Validasi input
	if name == "" || priceStr == "" || stockStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, price, and stock are required",
		})
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid price format",
		})
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid stock format",
		})
	}

	// Menangani Upload File
	file, err := c.FormFile("image")
	var imageURL string

	if err == nil {
		// Generate nama file yang unik untuk menghindari tabrakan nama file
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		savePath := filepath.Join("uploads", fileName)

		// Simpan file ke folder ./uploads
		if errSave := c.SaveFile(file, savePath); errSave != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan file gambar",
			})
		}
		// Set URL yang akan disimpan ke database (disajikan via fiber static)
		imageURL = fmt.Sprintf("/uploads/%s", fileName)
	}

	// Panggil usecase dengan tambahan URL gambar
	err = h.usecase.CreateProduct(c.Context(), name, description, price, stock, imageURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}
