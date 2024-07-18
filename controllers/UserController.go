package controllers

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"csprobe/server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	var users []models.User
	if result := inits.DATABASE.Find(&users); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error()})
	}
	context.JSON(http.StatusOK, users)
}

func GetUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User
	if err := inits.DATABASE.First(&user, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func PostUser(context *gin.Context) {
	var reqBody struct {
		Username uint
		Password string
	}
	context.Bind(&reqBody)
	hashedPassword, err := service.Encrypt(reqBody.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Hash the given password"})
		return
	}
	user := models.User{Username: reqBody.Username, Password: hashedPassword}
	if result := inits.DATABASE.Create(&user); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}
