package mysql

import (
	"diy-paas/utils"
	"fmt"
	"gorm.io/gorm/logger"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var db *gorm.DB

func Init() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
	)
	zapLogger, _ := zap.NewDevelopment()
	gormLogger := zapgorm2.New(zapLogger)
	gormLogger.SetAsDefault()
	gormLogger.LogLevel = logger.Info // 设置为 Info 级别

	if mysqlClient, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: gormLogger}); err != nil {
		utils.ErrorLog("mysql连接失败", "db", err.Error())
		panic(err)
	} else {
		zap.L().Info("mysql连接成功", zap.String("module", "db"))
		db = mysqlClient
		return mysqlClient
	}
}

func GetMysqlClient() *gorm.DB {
	return db
}

// 判断是否为重复条目错误
func IsDuplicateEntry(state string) bool {
	if strings.Contains(state, "Duplicate entry") {
		return true
	}
	return false
}
