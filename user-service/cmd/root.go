package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile menyimpan path ke file konfigurasi (jika diberikan via flag -c).
var cfgFile string

// rootCmd adalah perintah utama (root) dari aplikasi CLI kita.
// Kita menggunakan library Cobra untuk membuat Command Line Interface (CLI).
var rootCmd = &cobra.Command{
	Use:   "sayur-api",          // Nama perintah yang akan diketik di terminal.
	Short: "This API for sayur", // Penjelasan singkat perintah.
}

// Execute menjalankan perintah root dan menangani error jika ada.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// init adalah fungsi spesial di Go yang otomatis jalan sebelum main().
// Di sini kita mengatur flag (parameter) yang bisa diterima oleh CLI.
func init() {
	cobra.OnInitialize(initConfig) // Menghubungkan fungsi initConfig untuk dipanggil saat CLI inisialisasi.

	// Menambahkan flag global "--config" atau "-c" untuk menentukan file .env kustom.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")

	// Flag bawaan sebagai contoh bantuan (toggle).
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig bertugas untuk memuat file konfigurasi (.env) ke dalam sistem menggunakan Viper.
func initConfig() {
	if cfgFile != "" {
		// Jika user memberikan file spesifik via flag -c.
		viper.SetConfigFile(cfgFile)
	} else {
		// Jika tidak ada, default mengambil file ".env" di root folder.
		viper.SetConfigFile(".env")
	}

	// Otomatis membaca variabel environment yang ada di sistem operasi.
	viper.AutomaticEnv()

	// Mulai membaca isi file konfigurasi.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading config file:", err)
	}
}
