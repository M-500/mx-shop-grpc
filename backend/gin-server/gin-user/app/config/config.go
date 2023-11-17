package config

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 18:08
//

import (
	"gin-user/app/config/cfg"
	"github.com/spf13/viper"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:27
//

var ConfigInstance *Config

type Config struct {
	Name     string       `mapstructure:"name"`
	Host     string       `mapstructure:"host"`
	Local    string       `mapstructure:"local"`
	Port     int          `mapstructure:"port"`
	MySQl    cfg.MysqlCfg `mapstructure:"mysql"`
	Redis    cfg.RedisCfg `mapstructure:"redis"`
	Jwt      cfg.JwtCfg   `mapstructure:"jwt"`
	Logs     cfg.LogCfg   `mapstructure:"logs"`
	InitData cfg.BaseData `mapstructure:"base-data"`
}

func NewConfig(path string) *Config {
	ConfigInstance = &Config{}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(ConfigInstance); err != nil {
		panic(err)
	}
	return ConfigInstance
}
