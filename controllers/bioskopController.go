package controllers

import (
	"bioskop-api/config"
	"bioskop-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	err := config.DB.Select(&bioskops, "SELECT * FROM bioskop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	c.JSON(http.StatusOK, bioskops)
}

func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop
	err := config.DB.Get(&bioskop, "SELECT * FROM bioskop WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, bioskop)
}

func CreateBioskop(c *gin.Context) {
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, input.Nama, input.Lokasi, input.Rating).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat bioskop"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4`
	result, err := config.DB.Exec(query, input.Nama, input.Lokasi, input.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui bioskop"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil diperbarui"})
}

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	result, err := config.DB.Exec("DELETE FROM bioskop WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus bioskop"})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
