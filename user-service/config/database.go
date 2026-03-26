package config

import (
	"fmt"
	"user-service/database/seeds"
	"user-service/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql adalah wrapper struct untuk menahan instance database GORM.
type Mysql struct {
	DB *gorm.DB
}

// ConnectionMysql fungsi untuk membuat koneksi ke database MySQL menggunakan library GORM.
func ConnectionMysql(cfg *Config) (*Mysql, error) {
	// DSN format for MySQL: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	// Membuka koneksi database menggunakan GORM.
	db, err := gorm.Open(mysql.Open(dbConnString), &gorm.Config{})
	if err != nil {
		// Log error jika gagal konek dan mengembalikan error ke pemanggil fungsi.
		log.Error().Err(err).Msg("[ConnectionMysql-1] Failed to connect to database " + cfg.DB.Host)
		return nil, err
	}

	// Mengambil objek database SQL standar dari GORM untuk pengaturan lebih lanjut.
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionMysql-2] Failed to get database connection")
		return nil, err
	}

	// Automigrasi untuk membuat/menyesuaikan tabel database sesuai struct model.
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})

	// Menjalankan fungsi seeding untuk mengisi data awal (seperti Role dan Admin) secara otomatis.
	seeds.SeedRole(db)
	seeds.SeedAdmin(db)
	seeds.SeedUser(db)

	// Mengatur batas jumlah koneksi ke database agar lebih efisien.
	sqlDB.SetMaxOpenConns(cfg.DB.DBMaxOpen) // Batas maksimal koneksi yang aktif secara bersamaan.
	sqlDB.SetMaxIdleConns(cfg.DB.DBMaxIdle) // Batas maksimal koneksi yang tetap terbuka meski sedang tidak digunakan.

	// Mengembalikan pointer ke struct Mysql yang berisi instance database yang sudah siap dipakai.
	return &Mysql{DB: db}, nil
}
