package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB mengatur koneksi ke database MySQL
func ConnectDB() *gorm.DB {

	// Buat DSN untuk koneksi database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", Env.DBUser, Env.DBPass, Env.DBHost, Env.DBPort, Env.DBName)

	// Coba buka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Gagal terhubung ke MySQL: %v", err)
	}

	log.Println("Berhasil terhubung ke MySQL")

	return db

}
