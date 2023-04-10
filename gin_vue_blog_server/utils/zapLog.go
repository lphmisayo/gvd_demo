package utils

import (
	"gin_vue_blog_server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

func InitLogger() {
	if ok, _ := IsPathExist(global.Config.Zap.Directory); !ok {
		log.Printf("当前文件夹路径：%v 不存在，准备创建新文件夹\n", global.Config.Zap.Directory)
		err := os.Mkdir(global.Config.Zap.Directory, os.ModePerm)
		if err != nil {
			log.Println("创建文件夹失败！")
		} else {
			log.Println("创建文件夹成功！")
		}
	}
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), getLevelEnabler())
	global.Logger = zap.New(core)

	if global.Config.Zap.ShowLine {
		//获取 调用的文件、函数名称、行号
		global.Logger = global.Logger.WithOptions(zap.AddCaller())
	}

	log.Println("Zap Logger 初始化完成！")

}

// 编码器，用于初始化日志输入格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	if global.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	file, _ := os.Create(global.Config.Zap.Directory + "/log.log")

	//复数输出源
	if global.Config.Zap.LogInConsole {
		fileWriter := zapcore.AddSync(file)         //输出到文件
		consoleWriter := zapcore.AddSync(os.Stdout) //输出到命令行
		return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
	}

	//单数输出源
	return zapcore.AddSync(file)
}

func getLevelEnabler() zapcore.LevelEnabler {
	switch strings.ToLower(global.Config.Zap.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.Config.Zap.Prefix + t.Format("2006-01-02 - 15:04:05"))
}
