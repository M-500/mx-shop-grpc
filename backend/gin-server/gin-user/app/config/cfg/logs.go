// Package cfg
// Date           2023/2/15 21  28
// Version        1.0.0
// Author         代码小学生王木木
// Email          18574945291@163.com
// Description
package cfg

type LogCfg struct {
	Level         string `mapstructure:"level"`
	Path          string `mapstructure:"path"`
	Filename      string `mapstructure:"filename"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxAge        int    `mapstructure:"max_age"`
	MaxBackups    int    `mapstructure:"max_backups"`
	Compress      bool   `mapstructure:"compress"`
	StacktraceKey string `mapstructure:"stacktrace_key"`
	Format        string `mapstructure:"format"`
}
