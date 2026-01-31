package response

// DefaultResponse adalah struktur standar yang kita gunakan untuk semua respon API.
// Ini memastikan format JSON yang konsisten untuk client.
type DefaultResponse struct {
	Message string      `json:"message"`        // Pesan penjelasan hasil request (misal: "Success")
	Data    interface{} `json:"data,omitempty"` // Objek data hasil request. `omitempty` berarti field ini akan hilang jika kosong (nil).
}
