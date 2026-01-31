package cmd

import (
	"user-service/config"
	"user-service/internal/adapter/repository"
	"user-service/internal/app"
	"user-service/internal/core/service"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

// startCmd mendefinisikan perintah "start" yang digunakan untuk menjalankan server API.
// Contoh penggunaan di terminal: "go run main.go start"
var startCmd = &cobra.Command{
	Use:   "start",            // Nama sub-perintah.
	Short: "Start API server", // Deskripsi singkat.
	Long:  `Start API server`, // Deskripsi panjang.

	// Run adalah fungsi yang akan dijalankan ketika perintah "start" dipanggil.
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Inisialisasi Konfigurasi dari file .env.
		cfg := config.NewConfig()

		// 2. Membuat koneksi ke database PostgreSQL.
		db, err := config.ConnectionPostgres(cfg)
		if err != nil {
			// Jika gagal koneksi database, hentikan program aplikasi (Fatal).
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// 3. Dependency Injection (DI):
		// Kita menyusun objek (object) dari layer bawah ke atas.
		// Repository (akses data) -> Service (logika bisnis) -> Handler (HTTP request).
		userRepo := repository.NewUserRepository(db.DB) // Database dimasukkan ke Repository.
		userService := service.NewUserService(userRepo) // Repository dimasukkan ke Service.

		// 4. Menjalankan Server Utama dengan memasukkan Service yang sudah siap.
		app.RunServer(userService)
	},
}

func init() {
	// Mendaftarkan startCmd sebagai bagian dari perintah rootCmd.
	rootCmd.AddCommand(startCmd)
}
