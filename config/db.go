package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDB() {
	// Muat .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ File .env tidak ditemukan, menggunakan environment variable bawaan")
	}

	// Ambil variabel DATABASE_URL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL belum diset di environment atau file .env")
	}

	// Buka koneksi
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Gagal membuka koneksi:", err)
	}

	// Ping untuk cek koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}

	fmt.Println("✅ Koneksi database PostgreSQL berhasil")
}
