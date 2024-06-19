package main

import (
	"dmTool/core"
	"dmTool/global"
	"dmTool/info"
	"dmTool/server"
)

// 第三方包调用cmd，能实时输出
func main() {
	info.Info()
	// 初始化日志
	global.Log = core.InitLogger()
	// 读取配置文件
	core.InitConfig()
	//global.Log.Println("server info: ", global.Config)
	//global.Log.Println("db info: ", global.Config.Database.DbName)
	// 实例化导出导入的结构体对象
	server.ExpImp()
}
