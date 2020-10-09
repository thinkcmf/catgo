package main

import (
	_ "api_demo/api/admin"
	_ "api_demo/api/user"
	"github.com/thinkcmf/catgo"
	"github.com/gin-gonic/gin"
)

func init() {

	catgo.GET("/api", func(ctx *catgo.Context) {
		ctx.Success("恭喜您,API访问成功!", gin.H{
			"version": "1.0.0",
			"doc":     "http://www.thinkcmf.com/cmf5api.html",
		})
	})

	println("main router init")
}
