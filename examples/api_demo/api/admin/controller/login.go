package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo"
	"github.com/thinkcmf/catgo/examples/api_demo/api/admin/model"
	"strconv"
)

type LoginController struct{}

func (c *LoginController) Login(ctx *catgo.Context) {
	user := &model.User{}

	println(catgo.Db())
	catgo.DB.Debug().Find(user)

	println(strconv.FormatUint(user.Id, 10) + " " + user.UserNickname + " " + user.UserEmail)

	ctx.Success("登录成功！", gin.H{"user": user})
}
