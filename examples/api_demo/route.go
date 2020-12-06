package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo"
	_ "github.com/thinkcmf/catgo/examples/api_demo/api/admin"
	_ "github.com/thinkcmf/catgo/examples/api_demo/api/user"
)

func init() {

	catgo.DefaultGroup()
	{
		catgo.GET("/api", func(ctx *catgo.Context) {
			ctx.Success("恭喜您,API访问成功!", gin.H{
				"version": "1.0.0",
				"doc":     "http://www.thinkcmf.com/cmf5api.html",
			})
		})
		catgo.GET("/api2", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"version": "1.0.0",
				"doc":     "http://www.thinkcmf.com/cmf5api.html",
			})
		})
	}

	println("main router init")
}
