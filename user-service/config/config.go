package config

import "github.com/spf13/viper"

// App menyimpan konfigurasi dasar aplikasi seperti port dan environment (development/production).
type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
}

// JwtConfig menyimpan konfigurasi untuk JSON Web Token (JWT) yang digunakan untuk keamanan/autentikasi.
type JwtConfig struct {
	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

// PsqlDB menyimpan konfigurasi koneksi ke database PostgreSQL.
type PsqlDB struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	DBMaxOpen int    `json:"db_max_open"` // Batas maksimal koneksi yang terbuka.
	DBMaxIdle int    `json:"db_max_idle"` // Batas maksimal koneksi yang sedang tidak digunakan (idle).
}

// Config adalah struct utama yang menggabungkan semua konfigurasi (App, Jwt, Psql).
type Config struct {
	App  App       `json:"app"`
	Jwt  JwtConfig `json:"jwt"`
	Psql PsqlDB    `json:"psql"`
}

// NewConfig adalah fungsi generator untuk membuat instance Config baru dengan mengambil data dari environment variable menggunakan viper.
func NewConfig() *Config {
	return &Config{
		App: App{
			// Mengambil nilai port aplikasi dari variabel environment APP_PORT.
			AppPort: viper.GetString("APP_PORT"),
			// Mengambil nilai environment aplikasi (misal: development).
			AppEnv: viper.GetString("APP_ENV"),
		},
		Jwt: JwtConfig{
			// Mengambil secret key untuk enkripsi token JWT.
			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			// Mengambil nama issuer (pembuat) token JWT.
			JwtIssuer: viper.GetString("JWT_ISSUER"),
		},
		Psql: PsqlDB{
			// Mengambil konfigurasi detail database dari variabel environment.
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetString("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
	}
}
