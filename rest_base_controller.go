package catgo

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beecmf/beecmf/models"
)

type RestBaseController struct {
	beego.Controller
	currentUser *models.User
}

func (this *RestBaseController) Prepare() {
	//fmt.Println("prepare")
	//
	//user := this.CurrentUser()
	//
	//fmt.Println(user)
	//
	//if user == nil {
	//	this.Error("没有登录")
	//}

}

// 成功返回, 用法
// Sucsess("成功信息")
// Success("成功信息","跳转链接")
// Success("成功信息",返回数据struct)
// Success("成功信息",返回数据struct,"跳转链接")
func (this *RestBaseController) Success(message string, data ...interface{}) {

	fmt.Println("success")
	var (
		url        string
		mData      interface{}
		dataLength int
	)

	dataLength = len(data)

	if dataLength > 1 {

		mData = data[0]

		if v, ok := data[1].(string); ok {
			url = v
		}
	}

	if dataLength == 1 {
		if v, ok := data[0].(string); ok {
			url = v
		} else {
			mData = data[0]
		}

	}

	result := struct {
		Msg  string
		Url  string
		Data interface{}
		Code int
	}{
		Url:  url,
		Data: mData,
		Msg:  message,
		Code: 1,
	}

	this.Data["json"] = result

	this.setDevHeader()

	this.ServeJSON()

}

// 成功返回, 用法
// Error("错误信息")
// Error("错误信息","跳转链接")
// Error("错误信息",返回数据struct)
// Error("错误信息",返回数据struct,"跳转链接")
func (this *RestBaseController) Error(message string, data ...interface{}) {

	var (
		url        string
		mData      interface{}
		dataLength int
	)

	dataLength = len(data)

	if dataLength > 1 {

		mData = data[0]

		if v, ok := data[1].(string); ok {
			url = v
		}
	}

	if dataLength == 1 {
		if v, ok := data[0].(string); ok {
			url = v
		} else {
			mData = data[0]
		}

	}

	result := struct {
		Msg  string
		Url  string
		Data interface{}
		Code int
	}{
		Url:  url,
		Data: mData,
		Msg:  message,
		Code: 0,
	}

	this.Data["json"] = result

	this.setDevHeader()

	this.ServeJSON()

}

func (this *RestBaseController) setDevHeader() {
	if beego.BConfig.RunMode == beego.DEV {
		//this.Ctx.Output.Header("P3P", "CP=CAO PSA OUR")
		this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8010")
	}
}

func (this *RestBaseController) UserId() int64 {

	var userId int64 = 0
	sessionUserId := this.GetSession("UserId")

	fmt.Println(sessionUserId)

	switch v := sessionUserId.(type) {
	case int64:
		userId = int64(v)
	}

	return userId
}

func (this *RestBaseController) CurrentUser() *models.User {

	if this.currentUser == nil {
		var user models.User
		var userToken models.UserToken
		token := this.Ctx.Input.Header("XX-Token")
		if token != "" {
			hasToken, _ := Orm.Where("token = ?", token).Get(&userToken)

			if hasToken {
				hasUser, _ := Orm.Where("id = ?", userToken.Id).Get(&user)
				if hasUser {
					this.currentUser = &user
				}
			}
		}

	}

	return this.currentUser
}
