package config

import (
	"gin-vue-blog/utils/res"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}

	//Logger struct {
	//	Level        string
	//	Prefix       string
	//	Director     string
	//	ShowLine     bool // 是否显示行号
	//	LogInConsole bool //是否显示打印的路径
	//}

	System struct {
		Name string
		Port string
		Env  string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	//初始化数据库
	initDB()
	//初始化错误码
	res.ReadErrorCodeJson()
}
