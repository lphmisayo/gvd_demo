package config_model

type JWT struct {
	TokenKey string `yaml:"TokenKey" json:"token_key"`
	Expire   string `yaml:"Expire" json:"expire"`
	Issuer   string `yaml:"Issuer" json:"issuer"`
	Buffer   string `yaml:"Buffer" json:"buffer"`
}
