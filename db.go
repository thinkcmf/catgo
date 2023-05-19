package catgo

import (
	"fmt"
	"github.com/thinkcmf/catgo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB       *gorm.DB
	DBs      map[string]*gorm.DB
	DBConfig map[string]string
)

func Db(params ...interface{}) *gorm.DB {
	length := len(params)
	dbConfig := "default"
	if length > 0 {
		if value, ok := params[0].(string); ok {
			dbConfig = value
		}
	}

	var db *gorm.DB

	if _, ok := DBs[dbConfig]; ok {
		db = DBs[dbConfig]
	} else if dbConfig == "default" {
		dbInit()
		db = DBs[dbConfig]
	}

	return db
}

func Gorm(params ...interface{}) *gorm.DB {
	length := len(params)
	dbConfig := "default"
	if length > 0 {
		if value, ok := params[0].(string); ok {
			dbConfig = value
		}
	}

	var db *gorm.DB

	if _, ok := DBs[dbConfig]; ok {
		db = DBs[dbConfig]
	} else if dbConfig == "default" {
		dbInit()
		db = DBs[dbConfig]
	}

	return db
}

func dbInit() {

	dsn := ParseMainDbDsn()

	DBs = make(map[string]*gorm.DB)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DBConfig["prefix"], // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,               // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	DBs["default"] = DB

	if debug, ok := DBConfig["debug"]; ok {
		if debug == "true" {
			DBs["default"] = DB.Debug()
		}
	}
}

func ParseMainDbDsn() string {
	dbConfig := (config.New()).ReadDataBaseConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["database"],
		dbConfig["charset"],
	)

	DBConfig = dbConfig

	return dsn
}
