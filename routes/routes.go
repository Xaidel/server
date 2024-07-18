package routes

import (
	"csprobe/server/controllers"
	"csprobe/server/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
		}

		user := api.Group("/users")
		{
			user.GET("/", middleware.Authenticate, controllers.GetUsers)
			user.GET("/:id", controllers.GetUser)
			user.POST("/", controllers.PostUser)
		}
	}
}
