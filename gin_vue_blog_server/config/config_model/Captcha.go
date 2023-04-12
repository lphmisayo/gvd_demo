package config_model

type Captcha struct {
	SendEmail          bool `yaml:"SendEmail"`          // 是否通过邮箱发送验证码
	ExpireTime         int  `yaml:"ExpireTime"`         // 过期时间
	KeyLong            int  `yaml:"KeyLong"`            // 验证码长度
	ImgWidth           int  `yaml:"ImgWidth"`           // 验证码宽度
	ImgHeight          int  `yaml:"ImgHeight"`          // 验证码高度
	OpenCaptcha        int  `yaml:"OpenCaptcha"`        // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int  `yaml:"OpenCaptchaTimeOut"` // 防爆破验证码超时时间，单位：s(秒)
}
