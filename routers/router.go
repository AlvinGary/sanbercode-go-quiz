package routers

import (
	"database/sql"
	"sanbercode-go-quiz/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	kategoriGroup := router.Group("api/categories")
	{
		kategoriGroup.POST("", func(c *gin.Context) {
			controllers.CreateKategori(c, db)
		})
		kategoriGroup.GET("", func(c *gin.Context) {
			controllers.GetKategori(c, db)
		})
		kategoriGroup.GET("/:id", func(c *gin.Context) {
			controllers.GetKategoriByID(c, db)
		})
		kategoriGroup.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateKategori(c, db)
		})
		kategoriGroup.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteKategori(c, db)
		})
	}
}