package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-service/internal/adapter/handler"
	"user-service/internal/core/service"
	"user-service/utils/validator"

	"github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func RunServer(userService service.UserServiceInterface) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(
		customValidator.Validator,
		customValidator.Translator,
	)

	e.Validator = customValidator

	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	handler.NewUserHandler(e, userService)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := e.Start(":" + port); err != nil {
			log.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Print("[RunServer] Shutting down server in 5 seconds...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
