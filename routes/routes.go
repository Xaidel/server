package routes

import (
	"csprobe/server/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", controllers.GET)
		}
	}
}
