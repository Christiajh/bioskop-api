package controllers

import (
	"bioskop-api/config"
	"bioskop-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan bioskop"})
		return
	}

	c.JSON(http.StatusCreated, input)
}
