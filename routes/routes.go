package routes

import (
	"csprobe/server/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine){
	router.GET("/", controllers.GET)
}
