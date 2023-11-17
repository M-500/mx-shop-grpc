package config

import (
	"github.com/spf13/viper"
	"grpc-goods/config/cfg"
)

var ConfigInstance *Config

type Config struct {
	Name  string       `mapstructure:"name"`
	Host  string       `mapstructure:"host"`
	Local string       `mapstructure:"local"`
	Port  int          `mapstructure:"port"`
	MySQl cfg.MysqlCfg `mapstructure:"mysql"`
	Redis cfg.RedisCfg `mapstructure:"redis"`
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
