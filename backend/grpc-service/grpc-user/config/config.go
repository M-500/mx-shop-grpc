package config

import (
	"github.com/spf13/viper"
	"grpc-user/config/cfg"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:27
//

var ConfigInstance *Config

type Config struct {
	Name      string        `mapstructure:"name"`
	Host      string        `mapstructure:"host"`
	Local     string        `mapstructure:"local"`
	Port      int           `mapstructure:"port"`
	ConsulCfg cfg.ConsulCfg `mapstructure:"consul_cfg"`
	MySQl     cfg.MysqlCfg  `mapstructure:"mysql"`
	Redis     cfg.RedisCfg  `mapstructure:"redis"`
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
