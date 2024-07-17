package main
import ("github.com/gin-gonic/gin"
	"csprobe/server/inits"
	"csprobe/server/common"
	"csprobe/server/routes"
)

func init(){
	common.LoadEnv()
	inits.ConnectDB()
}

func main(){
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.InitRoutes(router)
	router.Run()
}
