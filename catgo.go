package catgo

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func Run() {
	initDbGorm()
}

func initDbGorm() {
	dbUser := beego.AppConfig.String("db.user")
	dbPass := beego.AppConfig.String("db.pass")
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.name")
	dbCharset := beego.AppConfig.String("db.charset")
	dbPrefix := beego.AppConfig.String("db.prefix")
	//dbDebug, _ := beego.AppConfig.Bool("db.debug")
	maxIdleConn, _ := beego.AppConfig.Int("db.max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db.max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbCharset) + "&loc=Asia%2FChongqing"
	fmt.Println(dbLink)
	
	var err error

	Db, err = gorm.Open("mysql", dbLink)

	if err != nil {
	println(err)
	}
	Db.DB().SetMaxIdleConns(maxIdleConn)
	Db.DB().SetMaxOpenConns(maxOpenConn)

	Db.SingularTable(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return dbPrefix + defaultTableName
	}
}