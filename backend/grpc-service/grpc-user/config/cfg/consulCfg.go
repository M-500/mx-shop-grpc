package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 19:41
//

type ConsulCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
