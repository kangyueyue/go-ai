package main

import (
	"fmt"
	"log"

	"github.com/kangyueyue/go-ai/common/mysql"
	"github.com/kangyueyue/go-ai/common/redis"
	"github.com/kangyueyue/go-ai/config"
	"github.com/kangyueyue/go-ai/router"
)

// StartServer 启动服务
func StartServer(addr string, port int) error {
	r := router.InitRouter()
	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

// main 主函数
func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	// init mysql
	if err := mysql.InitMysql(); err != nil {
		log.Println("InitMysql error , " + err.Error())
		return
	}
	// init redis
	redis.Init()

	// run
	err := StartServer(host, port)
	if err != nil {
		panic(err)
	}
}
