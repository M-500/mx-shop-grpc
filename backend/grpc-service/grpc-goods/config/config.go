package config

import "grpc-user/config/cfg"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:27
//
type Config struct {
	Name  string       `mapstructure:"name"`
	Host  string       `mapstructure:"host"`
	Local string       `mapstructure:"local"`
	Port  int          `mapstructure:"port"`
	MySQl cfg.MysqlCfg `mapstructure:"mysql"`
	Redis cfg.RedisCfg `mapstructure:"redis"`
}
