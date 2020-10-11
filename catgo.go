package catgo

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

var currentRouterGroup gin.IRouter

func init() {
	println("catgo init ")
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
	currentRouterGroup = Router
}

func Run() {
	println("catgo run ")
	Router.Run()
}
