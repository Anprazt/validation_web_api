package main

import (
	"net/http"

	"github.com/Anprazt/validation_web_api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Pong!",
		})
	})
	r.POST("/mahasiswa", handler.StudentHandler)
	r.Run(":8080")
}
