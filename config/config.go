package config

import (
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

	App struct {
		Name string
		Port int
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
	initDB()
}
