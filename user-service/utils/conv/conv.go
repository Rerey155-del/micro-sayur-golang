package conv

import "golang.org/x/crypto/bcrypt"

// HashPassword digunakan untuk mengubah password teks biasa (plain text)
// menjadi format acak (hash) agar aman saat disimpan di database.
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword mengenkripsi password dengan cost 14.
	// Cost menentukan seberapa kuat enkripsinya (semakin tinggi semakin aman tapi lambat).
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	// Mengembalikan hasil enkripsi dalam bentuk string.
	return string(bytes), nil
}

// CheckPasswordHash digunakan untuk mencocokkan password login (teks biasa)
// dengan hash password yang tersimpan di database.
func CheckPasswordHash(password, hash string) bool {
	// bcrypt.CompareHashAndPassword akan mengembalikan error nil jika password cocok.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
