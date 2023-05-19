package catgo

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (t *Config) ReadConfig(name string) (bool, error) {
	viper.SetConfigName(name)
	viper.AddConfigPath("./data/config")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Printf("Fatal error config file: %s \n", err)
	}
	return true, err
}

func (t *Config) IsSet(key string) bool {
	return viper.IsSet(key)
}

func (t *Config) Get(key string) interface{} {
	return viper.Get(key)
}

func (t *Config) GetString(key string) string {
	return viper.GetString(key)
}

func (t *Config) AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func (t *Config) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (t *Config) SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

func (t *Config) ReadDataBaseConfig() map[string]string {
	t.ReadConfig("database")
	t.SetDefault("user", "root")
	t.SetDefault("password", "")
	t.SetDefault("port", "3306")
	t.SetDefault("host", "mysql")
	t.SetDefault("database", "thinkcmf")
	t.SetDefault("charset", "utf8mb4")
	t.SetDefault("prefix", "cmf_")
	t.SetDefault("authcode", "")
	t.SetDefault("jwt_secret", "")
	t.SetDefault("debug", "false")
	user := t.GetString("user")
	password := t.GetString("password")
	host := t.GetString("host")
	port := t.GetString("port")
	database := t.GetString("database")
	charset := t.GetString("charset")
	authcode := t.GetString("authcode")
	jwt_secret := t.GetString("jwt_secret")
	debug := t.GetString("debug")
	return map[string]string{
		"user":       user,
		"password":   password,
		"host":       host,
		"port":       port,
		"database":   database,
		"charset":    charset,
		"authcode":   authcode,
		"jwt_secret": jwt_secret,
		"debug":      debug,
	}
}
