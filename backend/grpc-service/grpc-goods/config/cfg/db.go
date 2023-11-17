package cfg

type MysqlCfg struct {
	Datasource string `mapstructure:"datasource"`
}

type RedisCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"pass"`
}
