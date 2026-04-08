package model

// Product digunakan untuk memberitahu GORM bentuk tabel 'products' di database.
type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int     `gorm:"type:int;not null" json:"stock"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"`
}
