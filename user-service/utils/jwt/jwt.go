package jwt

import (
	"time"
	"user-service/internal/core/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims mendefinisikan payload data yang akan dimasukkan ke dalam token.
type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

// GenerateToken membuat token JWT baru berdasarkan data user.
func GenerateToken(user entity.UserEntity, secretKey string, issuer string) (string, error) {
	// Menentukan waktu kedaluwarsa token (misal: 24 jam ke depan).
	expirationTime := time.Now().Add(24 * time.Hour)

	// Membuat objek claims.
	claims := &JWTClaims{
		UserID:   user.ID,
		Email:    user.Email,
		RoleName: user.RoleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Membuat token menggunakan algoritma HS256 dan claims yang sudah dibuat.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token menggunakan Secret Key.
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
