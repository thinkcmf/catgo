package catgo

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strings"
)

var routes map[string]beego.ControllerInterface

func init() {
	routes = make(map[string]beego.ControllerInterface)
}

func AutoRouter(c beego.ControllerInterface) {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()
	pkgPath := t.PkgPath()

	apps := "/apps/"

	appsIndex := strings.LastIndex(pkgPath, apps)
	pkgPath = pkgPath[appsIndex:]

	app := strings.Replace(pkgPath, apps, "", -1)
	app = strings.Replace(app, "/controllers", "", -1)

	routes[app+"/"+t.Name()] = c

	fmt.Println(routes)

}
