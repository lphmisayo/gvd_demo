package core

import (
	"flag"
	"fmt"
	"gin_vue_blog_server/core/internal"
	"gin_vue_blog_server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "选择默认配置文件")
		flag.Parse()
		if config == "" { //若命令行输入为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您现在是%s模式,配置文件路径为%s\n", internal.ConfigEnv, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您现在是%s模式,配置文件路径为%s\n", internal.ConfigEnv, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您现在是%s模式,配置文件路径为%s\n", internal.ConfigEnv, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用命令行的-c参数传递的值,配置文件路径为%s\n", config)
			}
		} else {
			config = path[0]
			fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
		}
	}

	fmt.Println("当前配置路径:" + config)
	v := viper.New()
	//v.AddConfigPath(internal.ConfigDefaultPath)
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败，错误原因: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("配置文件被修改：", event.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println("配置文件解析式错误，错误原因：", err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println("配置文件解析式错误，错误原因：", err)
	}

	return v
}
