package routes

import (
	"csprobe/server/controllers"
	"csprobe/server/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

  controller := &controllers.Controller{}

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controller.Auth.Login)
		}

		user := api.Group("/users")
		{
			user.GET("/", middleware.Authenticate, controller.User.GetUsers)
			user.GET("/:id", controller.User.GetUser)
			user.POST("/", controller.User.PostUser)
			user.DELETE("/:id", controller.User.DeleteUser)
		}

    curriculum := api.Group("/curriculums")
    {
      curriculum.GET("/", middleware.Authenticate, controller.Curriculum.GET)
      curriculum.GET("/:id", middleware.Authenticate, controller.Curriculum.GET)
      curriculum.POST("/", middleware.Authenticate, controller.Curriculum.POST)
      curriculum.DELETE("/:id", middleware.Authenticate, controller.Curriculum.DELETE)
    }

	}
}
