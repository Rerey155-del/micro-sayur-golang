package model

import "time"

// User adalah struct model yang merepresentasikan tabel "users" di database PostgreSQL.
// GORM menggunakan tag `gorm` untuk menentukan konfigurasi kolom database.
type User struct {
	ID         int    `gorm:"primaryKey"` // Menandakan kolom ID sebagai Primary Key
	Name       string // Nama user (kolom string secara default)
	Email      string // Email user
	Password   string // Hash password user
	Address    string // Alamat user
	Phone      string // Nomor telepon user
	Photo      string // Path/Link foto user
	Lat        string // Koordinat Latitude
	Lng        string // Koordinat Longitude
	IsVerified bool   // Status verifikasi user

	// Metadata otomatis yang dikelola GORM (jika diaktifkan)
	CreatedAt time.Time  // Waktu data dibuat
	UpdatedAt time.Time  // Waktu data terakhir diupdate
	DeletedAt *time.Time // Waktu data dihapus (untuk Soft Delete)

	// Relasi Many-to-Many antara User dan Role melalui tabel perantara "user_roles"
	Roles []Role `gorm:"many2many:user_roles"`
}
