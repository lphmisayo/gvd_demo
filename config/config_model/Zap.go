package config_model

type Zap struct {
	Level        string `yaml:"Level"`        // 级别
	Prefix       string `yaml:"Prefix"`       // 日志前缀
	Format       string `yaml:"Format"`       // 输出
	Directory    string `yaml:"Directory"`    // 日志文件夹
	MaxAge       int    `yaml:"MaxAge"`       // 日志留存时间
	ShowLine     bool   `yaml:"ShowLine"`     // 显示行,显示文件位置
	LogInConsole bool   `yaml:"LogInConsole"` // 输出控制台
}
