package response

// SignInResponse adalah struktur data khusus yang dikembalikan setelah user berhasil login.
type SignInResponse struct {
	AccessToken string `json:"access_token"` // Token keamanan untuk akses API selanjutnya.
	Role        string `json:"role"`         // Nama role user.
	ID          int64  `json:"id"`           // ID unik user.
	Name        string `json:"name"`         // Nama user.
	Email       string `json:"email"`        // Email user.
	Lat         string `json:"lat"`          // Lokasi latitude user.
	Lng         string `json:"lng"`          // Lokasi longitude user.
	IsVerified  bool   `json:"is_verified"`  // Status verifikasi user.
}
