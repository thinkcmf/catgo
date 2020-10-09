package user

import (
	"github.com/thinkcmf/catgo"
	"github.com/thinkcmf/catgo/examples/api_demo/api/user/controller"
)

func init()  {
	var indexController *controller.IndexController
	catgo.GET("/api/user/index/index", indexController.Index)
	catgo.GET("/api/user/index/home", indexController.Home)
}