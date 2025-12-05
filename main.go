package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mylog "github.com/kangyueyue/go-ai/common/logger"
	"os"
	"strconv"

	"github.com/kangyueyue/go-ai/common/mysql"
	"github.com/kangyueyue/go-ai/common/redis"
	"github.com/kangyueyue/go-ai/config"
	"github.com/kangyueyue/go-ai/router"
)

// StartServer 启动服务
func StartServer(addr string, port int) error {
	r := router.InitRouter()
	mylog.Log.Infof("server start in port:%d", port)
	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

// main 主函数
func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	level := conf.MainConfig.Level
	appName := conf.MainConfig.AppName
	// init logger
	mylog.InitLog(level, appName)

	// init gin mode
	switch level {
	case "test":
		gin.SetMode(gin.TestMode) // 测试环境
	case "info":
		gin.SetMode(gin.ReleaseMode) // 线上环境
	default:
		gin.SetMode(gin.DebugMode) // 开发环境
	}

	// init mysql
	if err := mysql.InitMysql(); err != nil {
		mylog.Log.Errorf("InitMysql error %s", err.Error())
		return
	}
	// init redis
	redis.Init()

	// port param
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	err := StartServer(host, port)
	if err != nil {
		panic(err)
	}
}
