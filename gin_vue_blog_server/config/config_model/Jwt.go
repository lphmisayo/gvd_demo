package config_model

type JWT struct {
	Secret string `yaml:"Secret"`
	Expire string `yaml:"Expire"`
	Issuer string `yaml:"Issuer"`
}
