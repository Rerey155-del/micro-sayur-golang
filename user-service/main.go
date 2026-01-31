package main

import (
	"user-service/cmd"
)

// main adalah titik awal (entry point) eksekusi program Golang.
func main() {
	// Menjalankan fungsi Execute dari package cmd.
	// Fungsi ini akan mendeteksi sub-perintah yang diketik di terminal (seperti "start").
	cmd.Execute()
}
