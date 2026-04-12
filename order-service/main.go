package main

import (
	"log"
	"order-service/config"
	"order-service/internal/adapter/handler"
	"order-service/internal/adapter/repository"
	"order-service/internal/core/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.NewConfig()

	dbConn, err := config.ConnectionMysql(cfg)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	orderRepo := repository.NewOrderRepository(dbConn.DB)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	handler.NewOrderHandler(app, orderUsecase)

	log.Fatal(app.Listen(":" + cfg.App.AppPort))
}
