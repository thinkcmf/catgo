package controller

import (
	"api_demo/api/admin/model"
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo"
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
