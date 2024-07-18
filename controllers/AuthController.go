package controllers

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"csprobe/server/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(context *gin.Context) {
	// Define request body struct
	var reqBody struct {
		Username uint   `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := context.BindJSON(&reqBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := inits.DATABASE.Where("username = ?", reqBody.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})
		}
		return
	}

	if !service.Decrypt(reqBody.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}
