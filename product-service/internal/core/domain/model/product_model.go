package model

// Product digunakan untuk memberitahu GORM bentuk tabel 'products' di database.
type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"type:int;not null"`
}
