package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuefii/NusantaraKita/examples/api-go/database"
)

func main() {
	_, err := database.InitDB()
	if err != nil {
		panic("failed to connect database.")
	}
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "project is running"})
	})
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
