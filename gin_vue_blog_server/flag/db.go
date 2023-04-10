package flag

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/common"
)

func MakeMigration() {
	//系统邦定指令 - 映射数据库表
	fmt.Println("开始迁移数据库")
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&common.Article{},
			&common.User{},
			&common.Image{},
			&common.Comment{},
			&common.Tag{},
			&common.Category{},
		)
	if err != nil {
		global.Logger.Error(err.Error())
	}
}
