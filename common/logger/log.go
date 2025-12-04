package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"time"
)

// Log 日志
var Log *logrus.Logger

// InitLog 初始化日志
func InitLog(level, appName string) {
	Log = logrus.New()
	//Log.SetReportCaller(true) // 打印调用信息
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.DateTime,
		FullTimestamp:   true,
	})
	// level
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid logger level:%s", level))
	}
	// store file
	logFile := GetProjectPath() + "log/" + appName + ".log"
	out, err := rotatelogs.New(
		logFile+".%Y%m%d%H",                      // 指定日志文件的路径的名称，不存在自动创建
		rotatelogs.WithLinkName(logFile),         // 最新一份日志创建软链接,不带时间后缀
		rotatelogs.WithRotationTime(1*time.Hour), // 每隔1小时创建一份新的日志
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 只保留最近七天的日志，自动删除
	)
	if err != nil {
		panic(err)
	}
	Log.SetOutput(out)        // 设置日志文件
	Log.SetReportCaller(true) // 输出是从哪里调起的日志打印，代码行数
}

// GetProjectPath 获取项目路径
func GetProjectPath() string {
	_, currentPath, _, _ := runtime.Caller(0) // 当前行所在的文件目录
	return path.Dir(currentPath+"/../../../") + "/"
}
