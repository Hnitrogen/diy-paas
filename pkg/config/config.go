package config

import "github.com/spf13/viper"

func InitConfig(env string) {
	viper.SetConfigName(getConfigFileName(env))
	viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		panic("配置文件读取失败" + err.Error())
	}

}

func getConfigFileName(env string) string {
	switch env {
	case "dev":
		return "application.dev"
	case "prod":
		return "application.prod"
	default:
		return "application.dev"
	}
}
