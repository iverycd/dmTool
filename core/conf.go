package core

import (
	"dmTool/config"
	"dmTool/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig 初始化读取yaml配置信息并且解析到全局变量的结构体Config中
func InitConfig() {
	Log := logrus.New()
	viper.SetConfigFile("settings.yaml")
	// 通过viper读取配置文件进行加载
	if err := viper.ReadInConfig(); err == nil {
		Log.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		Log.Fatal(viper.ConfigFileUsed(), " has some error please check your yml file ! ", "Detail-> ", err)
	}
	conf := &config.Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		Log.Fatal(err)
	}
	global.Config = conf
}
