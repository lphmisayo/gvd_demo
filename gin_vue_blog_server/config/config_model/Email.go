package config_model

type Email struct {
	To       string // 收件人 多个以英文逗号分隔 例：a@qq.com,b@qq.com
	From     string // 发件人 要发邮件的邮箱
	Host     string `yaml:"Host"`     // 服务器地址, 例如 smtp.qq.com 前往要发邮件的邮箱查看其 smtp 协议
	Secret   string `yaml:"Secret"`   // 密钥, 不是邮箱登录密码, 是开启 smtp 服务后获取的一串验证码
	Nickname string `yaml:"Nickname"` // 发件人昵称, 通常为自己的邮箱名
	Port     int    `yaml:"Port"`     // 前往要发邮件的邮箱查看其 smtp 协议端口, 大多为 465
	IsSSL    bool   `yaml:"IsSSL"`    // 是否开启 SSL
}
