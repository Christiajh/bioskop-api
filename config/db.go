package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=postgres password=Sayabag dbname=silogydb sslmode=disable"

	var err error
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal membuka koneksi:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	fmt.Println("✅ Koneksi database PostgreSQL berhasil")
}
