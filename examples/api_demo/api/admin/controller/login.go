package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo"
	"github.com/thinkcmf/catgo/examples/api_demo/api/admin/model"
	"github.com/thinkcmf/catgo/examples/api_demo/api/admin/validate"
	"strconv"
)

type LoginController struct{}

func (c *LoginController) Login(ctx *catgo.Context) {
	user := &model.User{}

	loginForm := &validate.LoginForm{}
	if ctx.ShouldBind(loginForm) == nil {
		fmt.Println(loginForm.Password)
		fmt.Println(loginForm.Username)

		if result, msg := catgo.Validate(loginForm); !result {
			ctx.Error(msg)
			return
		}
	}

	println(catgo.Db())
	catgo.DB.Debug().Find(user)

	println(strconv.FormatUint(user.Id, 10) + " " + user.UserNickname + " " + user.UserEmail)

	ctx.Success("登录成功！", gin.H{"user": user})
}
