package admin

import (
	"github.com/thinkcmf/catgo/examples/api_demo/api/admin/controller"
	"github.com/thinkcmf/catgo"
)

func init() {
	var indexController *controller.LoginController
	catgo.POST("/api/admin/login", indexController.Login)
}
