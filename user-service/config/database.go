package config

import (
	"fmt"
	"user-service/database/seeds"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres adalah wrapper struct untuk menahan instance database GORM.
type Postgres struct {
	DB *gorm.DB
}

// ConnectionPostgres fungsi untuk membuat koneksi ke database PostgreSQL menggunakan library GORM.
// Fungsi ini juga menjalankan seeding data awal ke database.
func ConnectionPostgres(cfg *Config) (*Postgres, error) {
	// Membuat string koneksi (DSN - Data Source Name) dari konfigurasi yang ada.
	// sslmode=disable ditambahkan agar bisa konek ke database lokal tanpa memerlukan sertifikat SSL.
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.DBName)

	// Membuka koneksi database menggunakan GORM.
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		// Log error jika gagal konek dan mengembalikan error ke pemanggil fungsi.
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to database " + cfg.Psql.Host)
		return nil, err
	}

	// Mengambil objek database SQL standar dari GORM untuk pengaturan lebih lanjut.
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] Failed to get database connection")
		return nil, err
	}

	// Menjalankan fungsi seeding untuk mengisi data awal (seperti Role dan Admin) secara otomatis.
	seeds.SeedRole(db)
	seeds.SeedAdmin(db)

	// Mengatur batas jumlah koneksi ke database agar lebih efisien.
	sqlDB.SetMaxOpenConns(cfg.Psql.DBMaxOpen) // Batas maksimal koneksi yang aktif secara bersamaan.
	sqlDB.SetMaxIdleConns(cfg.Psql.DBMaxIdle) // Batas maksimal koneksi yang tetap terbuka meski sedang tidak digunakan.

	// Mengembalikan pointer ke struct Postgres yang berisi instance database yang sudah siap dipakai.
	return &Postgres{DB: db}, nil
}
