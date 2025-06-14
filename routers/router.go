package routers

import (
	"bioskop-api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Endpoint default root "/" agar tidak 404 saat akses domain langsung
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "✅ Bioskop API aktif. Gunakan endpoint /bioskop untuk mengambil data.",
		})
	})

	// Endpoint API bioskop
	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	return r
}
