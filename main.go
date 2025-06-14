package main

import (
	"bioskop-api/config"
	"bioskop-api/routers"
	"log"
)

func main() {
	// Hubungkan ke database
	config.ConnectDB()

	// Inisialisasi router
	r := routers.SetupRouter()

	// Jalankan server pada port 8000
	err := r.Run(":8000")
	if err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
