package catgo

import (
	"bytes"
	"catgo_demo/catgo/php"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB  *gorm.DB
	DBs map[string]*gorm.DB
)

func Db(db string) *gorm.DB {
	return DBs[db]
}

func init() {

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
		println(err.Error())
	}

	DBs["default"] = DB
}

func ParseMainDbDsn() string {
	currentPath, _ := GetCurrentPath()
	configFile := currentPath + "/conf/database.yml"

	databaseConfigContent, _ := php.FileGetContents(configFile)

	println(databaseConfigContent)

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

	log.Println(configMap)

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=" + charset + "&parseTime=True&loc=Local"

	log.Println(dsn)
	return dsn
}
