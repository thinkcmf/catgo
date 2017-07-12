package catgo

import (
	"github.com/astaxie/beego"
)

// this is BaseController
type BaseController struct {
	beego.Controller
}

func (this *BaseController) SetTheme(theme string) {

}

func (this *BaseController) Display(file string) {
	this.TplName = file + ".html"
}
