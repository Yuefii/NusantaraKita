package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuefii/NusantaraKita/examples/api-go/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	handler := handlers.NewHandler(db)

	router.GET("/api/provinces", handler.GetProvinces)
	router.GET("/api/regencies", handler.GetRegencies)
	router.GET("api/districts", handler.GetDistricts)
}
