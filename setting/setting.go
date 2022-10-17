package setting

import (
	"gopkg.in/ini.v1"
)

// app 配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

// 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

var Config = new(AppConfig)

func Init(file string) error {
	return ini.MapTo(Config, file)
}
