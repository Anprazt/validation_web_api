package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MahasiswaInput struct {
	NPM    int    `json:"NPM" binding:"required"`
	Nama   string `json:"Nama" binding:"required"`
	Email  string `json:"Email" binding:"required,email"`
	Prodi  string `json:"Prodi" binding:"required"`
	Alamat string `json:"Alamat" binding:"required"`
}

func StudentHandler(c *gin.Context) {
	var mahasiswaInput MahasiswaInput
	time := time.Now()

	err := c.ShouldBindJSON(&mahasiswaInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append(errorMessages, report)
			case "email":
				report := fmt.Sprintf("%s is not valid email", e.Field())
				errorMessages = append(errorMessages, report)
			}
			//errorMessage := fmt.Sprintf("Field %s harus diisi. Pesan: %s", e.Field(), e.ActualTag())
			//errorMessages = append(errorMessages, report)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Berhasil input data",
		"NPM":     mahasiswaInput.NPM,
		"Nama":    mahasiswaInput.Nama,
		"Email":   mahasiswaInput.Email,
		"Prodi":   mahasiswaInput.Prodi,
		"Alamat":  mahasiswaInput.Alamat,
		"Time":    time,
	})
}
