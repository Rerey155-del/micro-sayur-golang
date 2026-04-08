package main

import (
	"log"
	"product-service/config"
	"product-service/internal/adapter/handler"
	"product-service/internal/adapter/repository"
	"product-service/internal/core/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Load Konfigurasi
	cfg := config.NewConfig()

	// 2. Koneksi Database
	dbConn, err := config.ConnectionMysql(cfg)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// 3. Inisialisasi Repository, Usecase, dan Handler
	productRepo := repository.NewProductRepository(dbConn.DB)
	productUsecase := usecase.NewProductUsecase(productRepo)

	// 4. Setup Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New()) // Aktifkan CORS agar Frontend bisa memanggil API ini

	// Menyajikan folder uploads agar bisa diakses langsung via HTTP
	app.Static("/uploads", "./uploads")

	// 5. Register Routes
	handler.NewProductHandler(app, productUsecase)

	// 6. Jalankan Server
	log.Fatal(app.Listen(":" + cfg.App.AppPort))
}
