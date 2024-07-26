package controllers

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"csprobe/server/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct{}

func (auth *AuthController) Login(context *gin.Context) {
	var reqBody struct {
		Username uint
		Password string
	}

	if context.Bind(&reqBody) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var user models.User

	inits.DATABASE.First(&user, "username = ?", reqBody.Username)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username",
		})
		return
	}

	if !service.Decrypt(reqBody.Password, user.Password) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate a token",
		})
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{})

}
