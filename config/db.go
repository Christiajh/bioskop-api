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
	// Muat .env file (opsional di local dev)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ File .env tidak ditemukan, menggunakan environment variable bawaan")
	}

	// Ambil DATABASE_URL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL belum diset di environment atau file .env")
	}

	// Buka koneksi
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Gagal membuka koneksi:", err)
	}

	// Ping
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database:", err)
	}

	fmt.Println("✅ Koneksi database PostgreSQL berhasil")

	// 👇 Buat tabel bioskop jika belum ada
	createTable := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		lokasi VARCHAR(100) NOT NULL,
		rating REAL
	);
	`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("❌ Gagal membuat tabel bioskop:", err)
	}
	fmt.Println("✅ Tabel bioskop sudah tersedia atau berhasil dibuat")
}
