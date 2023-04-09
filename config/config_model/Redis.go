package config_model

type Redis struct {
	DB       int    `yaml:"DB"`       //指定 Redis 数据库
	Addr     string `yaml:"Addr"`     // 服务器地址:端口
	Password string `yaml:"Password"` // 密码
}
