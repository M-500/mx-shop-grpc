// Package cfg
// Date        : 2023/2/15 21:22
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package cfg

type MysqlCfg struct {
	Datasource string `mapstructure:"datasource"`
}

type RedisCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"pass"`
}
