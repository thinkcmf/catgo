package catgo

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	xormCore "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Orm *xorm.Engine

func Run() {
	initView()
	//initDb()
	initDbXorm()
	initValidation()
	beego.Run()
}

func initView() {
	beego.SetViewsPath("public/themes/default/")
}

func initDbXorm() {
	dbUser := beego.AppConfig.String("db.user")
	dbPass := beego.AppConfig.String("db.pass")
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.name")
	dbCharset := beego.AppConfig.String("db.charset")
	dbPrefix := beego.AppConfig.String("db.prefix")
	dbDebug, _ := beego.AppConfig.Bool("db.debug")
	maxIdleConn, _ := beego.AppConfig.Int("db.max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db.max_open_conn")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbCharset) + "&loc=Asia%2FChongqing"
	fmt.Println(dbLink)
	var err error
	Orm, err = xorm.NewEngine("mysql", dbLink)
	Orm.ShowSQL(dbDebug)
	Orm.SetMaxOpenConns(maxOpenConn)
	Orm.SetMaxIdleConns(maxIdleConn)
	tbMapper := xormCore.NewPrefixMapper(xormCore.SnakeMapper{}, dbPrefix)
	Orm.SetTableMapper(tbMapper)

	fmt.Println(err)
}

func initDb() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// database
	dbUser := beego.AppConfig.String("db.user")
	dbPass := beego.AppConfig.String("db.pass")
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.name")
	dbCharset := beego.AppConfig.String("db.charset")
	maxIdleConn, _ := beego.AppConfig.Int("db.max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db.max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbCharset) + "&loc=Asia%2FChongqing"
	// set default database
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)
}

func initValidation() {
	//设置表单验证提示信息
	messageTmpls := map[string]string{
		"Required":     "Can not be empty 1",
		"Min":          "Minimum is %d",
		"Max":          "Maximum is %d",
		"Range":        "Range is %d to %d",
		"MinSize":      "Minimum size is %d",
		"MaxSize":      "Maximum size is %d",
		"Length":       "Required length is %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	}

	validation.SetDefaultMessage(messageTmpls)
}
