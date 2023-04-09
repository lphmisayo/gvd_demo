package config_model

type Mysql struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Config   string `yaml:"Config"`
	Db       string `yaml:"Db"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	LogLevel string `yaml:"Log_Level"`
}
