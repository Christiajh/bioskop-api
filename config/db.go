package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=postgres password=Sayakaje dbname=terlalu sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	fmt.Println("Koneksi database berhasil")
}
