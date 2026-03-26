package config

import (
	"fmt"
	"product-service/database/seeds"
	"product-service/internal/core/domain/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func ConnectionMysql(cfg *Config) (*Mysql, error) {
	// DSN format for MySQL: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate Tabel
	db.AutoMigrate(&model.Product{})

	// Menjalankan Seeding Data
	seeds.SeedProduct(db)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.DB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.DB.DBMaxIdle)

	return &Mysql{DB: db}, nil
}
