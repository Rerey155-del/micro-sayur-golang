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

// RunServer adalah jantung dari aplikasi kita. Fungsi ini mengatur setting HTTP server (Echo),
// validasi, endpoint, dan mekanisme shutdown server agar aman (Graceful Shutdown).
func RunServer(userService service.UserServiceInterface) {
	// 1. Membuat instance baru dari web framework Echo.
	e := echo.New()

	// 2. Mendaftarkan Middleware. Middleware adalah kode yang jalan sebelum Handler dipanggil.
	e.Use(middleware.CORS())    // Mengizinkan request dari domain lain (Cross-Origin Resource Sharing).
	e.Use(middleware.Recover()) // Mencegah server mati jika ada panic error.

	// 3. Konfigurasi Validator untuk validasi input (misal: format email harus valid).
	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(
		customValidator.Validator,
		customValidator.Translator,
	)
	e.Validator = customValidator

	// 4. Endpoint sederhana untuk cek status server secara manual.
	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// 5. Mendaftarkan User Handler ke instance Echo agar endpoint /signin bisa diakses.
	handler.NewUserHandler(e, userService)

	// 6. Mengambil port dari environment variable (misal: 8080).
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// 7. Menjalankan server di background (goroutine) supaya kode selanjutnya tetap bisa jalan.
	go func() {
		if err := e.Start(":" + port); err != nil {
			log.Error(err)
		}
	}()

	// 8. Mekanisme Graceful Shutdown:
	// Baris ini akan "menunggu" sampai ada sinyal tutup dari sistem (seperti menekan Ctrl+C).
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Program berhenti di sini sampai menerima sinyal interupsi.

	log.Print("[RunServer] Shutting down server in 5 seconds...")

	// 9. Memberi waktu 5 detik bagi koneksi yang masih aktif untuk selesai sebelum benar-benar mati.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
