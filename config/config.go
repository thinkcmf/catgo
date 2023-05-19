package config

import (
	"github.com/thinkcmf/catgo"
)

type Config struct {
}

func New() *Config {
	return &Config{}
}

func (t *Config) ReadDataBaseConfig() map[string]string {
	config := catgo.NewConfig()
	config.ReadConfig("database")
	config.SetDefault("user", "root")
	config.SetDefault("password", "")
	config.SetDefault("port", "3306")
	config.SetDefault("host", "mysql")
	config.SetDefault("database", "thinkcmf")
	config.SetDefault("charset", "utf8mb4")
	config.SetDefault("prefix", "cmf_")
	config.SetDefault("authcode", "")
	config.SetDefault("jwt_secret", "")
	config.SetDefault("debug", "false")
	user := config.GetString("user")
	password := config.GetString("password")
	host := config.GetString("host")
	port := config.GetString("port")
	database := config.GetString("database")
	charset := config.GetString("charset")
	authcode := config.GetString("authcode")
	jwt_secret := config.GetString("jwt_secret")
	debug := config.GetString("debug")
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
