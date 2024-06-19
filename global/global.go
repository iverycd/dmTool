package global

import (
	"dmTool/config"
	"github.com/sirupsen/logrus"
)

var (
	Config *config.Config // 是config包内Config结构体,从这里获取配置文件中各个信息
	Log    *logrus.Logger
)
