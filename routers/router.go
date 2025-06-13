package routers

import (
	"bioskop-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	return r
}
