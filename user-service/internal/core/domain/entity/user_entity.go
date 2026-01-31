package entity

// UserEntity adalah struktur data utama yang mewakili User dalam domain aplikasi kita.
// Entity digunakan untuk mentransfer data antar layer (Handler <-> Service <-> Repository)
// tanpa harus bergantung pada struktur tabel database (model).
type UserEntity struct {
	ID         int64  // ID unik user
	Name       string // Nama lengkap user
	Email      string // Alamat email user
	Password   string // Hash password user
	Address    string // Alamat tempat tinggal user
	RoleName   string // Nama role user (misal: Customer atau Super Admin)
	Phone      string // Nomor telepon user
	Photo      string // Link/path foto profil user
	Lat        string // Koordinat Latitude (lokasi)
	Lng        string // Koordinat Longitude (lokasi)
	IsVerified bool   // Status apakah user sudah diverifikasi atau belum
}
