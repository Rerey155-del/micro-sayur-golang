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
	api.Put("/products/:id", handler.UpdateProduct)
	api.Delete("/products/:id", handler.DeleteProduct)
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

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	existingImageURL := c.FormValue("image_url") // To keep old image if no new one uploaded

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

	// Menangani Upload File Baru (Opsional)
	file, err := c.FormFile("image")
	imageURL := existingImageURL

	if err == nil {
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		savePath := filepath.Join("uploads", fileName)
		if errSave := c.SaveFile(file, savePath); errSave != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan file gambar baru",
			})
		}
		imageURL = fmt.Sprintf("/uploads/%s", fileName)
	}

	err = h.usecase.UpdateProduct(c.Context(), uint(id), name, description, price, stock, imageURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	err = h.usecase.DeleteProduct(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
