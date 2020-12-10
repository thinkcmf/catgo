package catgo

import (
	"bytes"
	"fmt"
	"github.com/icatgo/php"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	DBs map[string]*gorm.DB
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

func dbInit() {

	dsn := ParseMainDbDsn()

	DBs = make(map[string]*gorm.DB)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cmf_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	DBs["default"] = DB
}

func ParseMainDbDsn() string {
	configFile := "./conf/database.yml"

	databaseConfigContent, _ := php.FileGetContents(configFile)

	data := []byte(databaseConfigContent)
	yamlDecoder := yaml.NewDecoder(bytes.NewBuffer(data))

	var configMap map[string]string
	err2 := yamlDecoder.Decode(&configMap)

	if err2 != nil {
		println(err2.Error())
	}

	user := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	database := ""
	charset := "utf8mb4"

	for key, value := range configMap {
		switch key {
		case "user":
			user = value
		case "password":
			password = value
		case "host":
			host = value
		case "port":
			port = value
		case "database":
			database = value
		case "charset":
			charset = value
		}
	}

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=" + charset + "&parseTime=True&loc=Local"

	return dsn
}
