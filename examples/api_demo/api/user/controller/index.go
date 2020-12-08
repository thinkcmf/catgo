package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo"
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

	fmt.Println(catgo.DB)
	catgo.Db().Debug().Find(user)

	println(strconv.FormatInt(user.ID, 10) + " " + user.UserNickname + " " + user.UserEmail)
	ctx.JSON(200, gin.H{
		"msg": "hello",
	})
}

func (c *IndexController) Home(ctx *catgo.Context) {
	ctx.Success("success", "ddd")
}
