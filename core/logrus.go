package core

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() *logrus.Logger {
	mLog := logrus.New()        //新建一个实例
	mLog.SetOutput(os.Stdout)   //设置输出类型
	mLog.SetReportCaller(false) //开启返回函数名和行号
	return mLog
}
