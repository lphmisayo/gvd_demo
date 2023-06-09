package utils

import (
	"gin_vue_blog_server/config"
	"gin_vue_blog_server/global"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	ConfigFile = "config/config.yaml"
)

func InitConfig() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Panic("配置文件读取失败: ", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Panic("配置文件解析失败：", err)
	}
	log.Println("配置文件解析完成")
	global.Config = c
}
