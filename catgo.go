package catgo

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	println("catgo init ")
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
}

func Run() {
	println("catgo run ")
	Router.Run()
}


