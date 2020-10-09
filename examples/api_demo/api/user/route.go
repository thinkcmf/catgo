package user

import (
	"api_demo/api/user/controller"
	"github.com/thinkcmf/catgo"
)

func init()  {
	var indexController *controller.IndexController
	catgo.GET("/api/user/index/index", indexController.Index)
	catgo.GET("/api/user/index/home", indexController.Home)
}