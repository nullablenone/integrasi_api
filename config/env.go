package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ENV adalah struct untuk menyimpan konfigurasi environment
type ENV struct {
	DBUser         string
	DBPass         string
	DBHost         string
	DBPort         string
	DBName         string

	ExternalAPIURL string
}

// Variabel global untuk menyimpan konfigurasi environment yang sudah dimuat
var Env ENV

// LoadENV digunakan untuk memuat variabel environment dari file .env
func LoadENV() {

	// Memuat variabel dari file .env
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error saat memuat file .env: %v", err)
	}

	// Isi struct Env dengan nilai dari environment variable
	Env = ENV{
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBName:         os.Getenv("DB_NAME"),
		
		ExternalAPIURL: os.Getenv("EXTERNAL_API_URL"),
	}
	
}
