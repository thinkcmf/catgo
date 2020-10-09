package controller

import (
	"github.com/thinkcmf/catgo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IndexController struct {
}

type User struct {
	ID           int64
	UserNickname string
	UserEmail    string
}

func (c *IndexController) Index(ctx *catgo.Context) {
	user := &User{}

	println(catgo.DB)
	catgo.DB.Debug().Find(user)

	println(strconv.FormatInt(user.ID, 10) + " " + user.UserNickname + " " + user.UserEmail)
	ctx.JSON(200, gin.H{
		"msg": "hello",
	})
}

func (c *IndexController) Home(ctx *catgo.Context) {
	ctx.Success("success", "ddd")
}
