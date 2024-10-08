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
			user.GET("/", middleware.Authenticate, controller.User.GET)
			user.GET("/:id", middleware.Authenticate, controller.User.GET)
			user.POST("/", controller.User.POST)
			user.DELETE("/:id", middleware.Authenticate, controller.User.DELETE)
		}

		curriculum := api.Group("/curriculums")
		{
			curriculum.GET("/", middleware.Authenticate, controller.Curriculum.GET)
			curriculum.GET("/:id", middleware.Authenticate, controller.Curriculum.GET)
			curriculum.POST("/", middleware.Authenticate, controller.Curriculum.POST)
			curriculum.DELETE("/:id", middleware.Authenticate, controller.Curriculum.DELETE)
		}

		course := api.Group("/courses")
		{
			course.GET("/", middleware.Authenticate, controller.Course.GET)
			course.GET("/:id", middleware.Authenticate, controller.Course.GET)
			course.POST("/", middleware.Authenticate, controller.Course.POST)
			course.DELETE("/:id", middleware.Authenticate, controller.Course.DELETE)
		}

	}
}
