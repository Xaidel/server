package controllers
import ("github.com/gin-gonic/gin")

func GET(context *gin.Context){
		context.JSON(200, gin.H{
			"message": "user get endpoint",
		})
}
