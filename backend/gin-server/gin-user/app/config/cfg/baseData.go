//@Author: wulinlin
//@Description:
//@File:  baseData
//@Version: 1.0.0
//@Date: 2023/03/15 14:38

package cfg

type BaseData struct {
	AdminUser AdminInfo `mapstructure:"admin-user"`
	CateList  []string  `mapstructure:"cate-list"`
}

type AdminInfo struct {
	Username string `mapstructure:"username"`
	Phone    string `mapstructure:"phone"`
	Password string `mapstructure:"password"`
}
