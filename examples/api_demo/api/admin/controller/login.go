package controller

import (
	"github.com/thinkcmf/catgo/examples/api_demo/api/admin/model"
	"github.com/thinkcmf/catgo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type LoginController struct{}

func (c *LoginController) Login(ctx *catgo.Context) {
	user := &model.User{}

	println(catgo.DB)
	catgo.DB.Debug().Find(user)

	println(strconv.FormatUint(user.Id, 10) + " " + user.UserNickname + " " + user.UserEmail)

	ctx.Success("登录成功！", gin.H{"user": user})
}
