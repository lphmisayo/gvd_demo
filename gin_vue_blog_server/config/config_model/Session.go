package config_model

type Session struct {
	Name   string `yaml:"Name"`
	Salt   string `yaml:"Salt"`
	MaxAge int    `yaml:"MaxAge"`
}
