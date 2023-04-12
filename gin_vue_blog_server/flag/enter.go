package flag

import (
	sys_flag "flag"
	confM "gin_vue_blog_server/config/config_model"
)

func Parse() *confM.Option {

	var config string

	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.StringVar(&config, "c", "", "选择默认配置文件")

	sys_flag.Parse()
	return &confM.Option{
		DB:     *db,
		Config: config,
	}
}

func IsWebStop(option *confM.Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option *confM.Option) {
	if option.DB {
		MakeMigration()
	}
}
