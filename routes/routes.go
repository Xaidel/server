package routes

import (
	"csprobe/server/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.GET("/", controllers.GetUsers)
			user.GET("/:id", controllers.GetUser)
			user.POST("/", controllers.PostUser)
		}
	}
}
