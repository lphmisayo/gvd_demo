package initialize

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

func InitMysqlDB() *gorm.DB {
	mysqlDB := global.Config.Mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlDB.User,
		mysqlDB.Password,
		mysqlDB.Host,
		mysqlDB.Port,
		mysqlDB.Db,
	)

	db, err := gorm.Open(mysql.Open(dns), getGormConfig())
	if err != nil {
		log.Fatal("Mysql连接失败，请检查参数")
	}
	log.Println("Mysql连接成功！")

	//迁移数据库表 绑定至console语句使用

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)                  //设置连接池中的最大闲连接
	sqlDB.SetMaxOpenConns(100)                 //设置数据库的最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) //设置连接的最大可复用时间

	return db
}

func getGormConfig() *gorm.Config {
	return &gorm.Config{
		//设置日志模式
		Logger: logger.Default.LogMode(getLogMode(global.Config.Mysql.LogLevel)),
		//禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		//禁用默认事务（提升运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
	}
}

func getLogMode(level string) logger.LogLevel {
	switch strings.ToLower(level) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}
