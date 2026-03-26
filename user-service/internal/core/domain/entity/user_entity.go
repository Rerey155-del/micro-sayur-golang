package entity

// UserEntity adalah struktur data utama yang mewakili User dalam domain aplikasi kita.
// Entity digunakan untuk mentransfer data antar layer (Handler <-> Service <-> Repository)
// tanpa harus bergantung pada struktur tabel database (model).
type UserEntity struct {
	ID         int64  `json:"id"`          // ID unik user
	Name       string `json:"name"`        // Nama lengkap user
	Email      string `json:"email"`       // Alamat email user
	Password   string `json:"password"`    // Hash password user
	Address    string `json:"address"`     // Alamat tempat tinggal user
	RoleName   string `json:"role_name"`   // Nama role user (misal: Customer atau Super Admin)
	Phone      string `json:"phone"`       // Nomor telepon user
	Photo      string `json:"photo"`       // Link/path foto profil user
	Lat        string `json:"lat"`         // Koordinat Latitude (lokasi)
	Lng        string `json:"lng"`         // Koordinat Longitude (lokasi)
	IsVerified bool   `json:"is_verified"` // Status apakah user sudah diverifikasi atau belum
}
