package initialize

import (
	"diy-paas/internal/app/model"
	"diy-paas/internal/global"
)

func InitTables() {
	global.Mysql.AutoMigrate(&model.User{}) // 用户表
}
