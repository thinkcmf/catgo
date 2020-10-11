package catgo

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
	currentRouterGroup gin.IRouter
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
	currentRouterGroup = Router
}

func Run(addr ...string) {
	Router.Run(addr...)
}
