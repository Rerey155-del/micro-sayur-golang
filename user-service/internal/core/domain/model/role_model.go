package model

import "time"

// Role merepresentasikan tabel "roles" di database (seperti Super Admin atau Customer).
type Role struct {
	ID   int64  `gorm:"primaryKey"` // ID unik role
	Name string // Nama role

	// Menghubungkan Role kembali ke User (Many-to-Many)
	Users []User `gorm:"many2many:user_roles"`

	CreatedAt time.Time  // Waktu data role dibuat
	UpdatedAt time.Time  // Waktu data role diupdate
	DeletedAt *time.Time // Waktu data role dihapus (Soft Delete)
}
