package config

import "gin_vue_blog_demo/config/config_model"

type Config struct {
	Server  config_model.Server  `yaml:"Server"`
	Mysql   config_model.Mysql   `yaml:"Mysql"`
	JWT     config_model.JWT     `yaml:"JWT"`
	Redis   config_model.Redis   `yaml:"Redis"`
	Session config_model.Session `yaml:"Session"`
	Upload  config_model.Upload  `yaml:"Upload"`
	Email   config_model.Email   `yaml:"Email"`
	Captcha config_model.Captcha `yaml:"Captcha"`
	Zap     config_model.Zap     `yaml:"Zap"`
}
