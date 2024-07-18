package middleware

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(context *gin.Context) {
	tokenString, err := context.Cookie("Authorization")

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		inits.DATABASE.First(&user, claims["sub"])

		if user.ID == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		context.Set("user", user)

		context.Next()

	} else {
		fmt.Println(err)
	}

}
