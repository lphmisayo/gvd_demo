package config_model

type Captcha struct {
	SendEmail  bool `yaml:"SendEmail"`  // 是否通过邮箱发送验证码
	ExpireTime int  `yaml:"ExpireTime"` // 过期时间
}
