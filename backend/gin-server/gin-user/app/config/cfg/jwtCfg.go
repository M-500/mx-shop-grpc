// Package cfg
// Date        : 2023/2/15 21:22
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package cfg

type JwtCfg struct {
	SecretKey     string `mapstructure:"secret_key"`
	JwtHeaderKey  string `mapstructure:"jwt_header_key"`
	JwtPrefix     string `mapstructure:"jwt_prefix"`
	JwtExpireHour int64  `mapstructure:"jwt_expire_hour"`
}
