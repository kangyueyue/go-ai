package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai/common/aihelper"
	"github.com/kangyueyue/go-ai/common/logger"
	mylog "github.com/kangyueyue/go-ai/common/logger"
	"github.com/kangyueyue/go-ai/common/mq"
	"github.com/kangyueyue/go-ai/dao/message"

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
	// 初始化aihelper
	readDataFromDB()

	// init redis
	redis.Init()
	mylog.Log.Infof("redis init success")

	// init mq
	mq.InitRabbitMq()
	logger.Log.Infof("mq init success")

	// port param
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	// start http server
	err := StartServer(host, port)
	if err != nil {
		panic(err)
	}
}

// 从数据库加载消息并初始化 AIHelperManager
func readDataFromDB() error {
	manager := aihelper.GetGlobalManager()
	msgs, err := message.GetAllMessages()
	if err != nil {
		return err
	}
	// for
	for i := range msgs {
		m := &msgs[i]
		// 默认openai 模型
		modelType := "1"
		config := make(map[string]interface{})

		// 创建对应的AIHelper
		helper, err := manager.GetOrCreateAIHelper(m.UserName, m.SessionID, modelType, config)
		if err != nil {
			logger.Log.Infof("[readDataFromDB] failed to create helper for user=%s session=%s: %v", m.UserName, m.SessionID, err)
			continue
		}
		logger.Log.Info("readDataFromDB init:  ", helper.SessionID)
		helper.AddMessage(m.Content, m.UserName, m.IsUser, false)
	}
	logger.Log.Infof("AIHelperManager init success ")
	return nil
}
