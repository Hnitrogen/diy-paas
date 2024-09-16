package main

import (
	"diy-paas/internal/global"
	"diy-paas/internal/initialize"
	"diy-paas/pkg/config"
	"diy-paas/pkg/logger"
	"diy-paas/pkg/mysql"
	"diy-paas/pkg/server"
	"flag"
	"fmt"
)

func main() {
	/*
		BaseApp
			主要负责登录注册等站点基础功能
			基础配置初始化
	*/
	fmt.Println("启动 BaseApp......")
	env := flag.String("env", "dev", "dev/prod")
	//	C:\Users\17914\GolandProjects\diy-paas\conf\application.dev.ini
	flag.Parse()

	// 初始化配置文件
	config.InitConfig(*env)
	// 初始化日志组件
	logger.InitLogger()
	// 初始化表结构
	global.Mysql = mysql.Init()
	initialize.InitTables()
	// 启动rest-ful server
	server.InitRestFulServer()
}
