package main

import (
	"github.com/thinkcmf/catgo"
	"github.com/thinkcmf/catgo/im"
	"github.com/astaxie/beego"
)

func main() {
	// Register routers.
	//beego.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	//beego.Router("/join", &controllers.AppController{}, "post:Join")

	beego.Router("/ws/client", &im.WebSocketController{}, "get:Client")
	beego.Router("/ws/client", &im.WebSocketController{}, "post:Client")
	beego.Router("/ws/join", &im.WebSocketController{}, "get:Join")
	catgo.Run()
}
