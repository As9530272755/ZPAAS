package base

import (
	"github.com/astaxie/beego"
)

func Log() {
	// 开启日志
	beego.BConfig.Log.AccessLogs = true

	// 日志里面显示文件名和输出日志行号
	beego.BConfig.Log.FileLineNum = true

	// 不在终端打印日志
	beego.BeeLogger.DelLogger("console")

	// 日志输出配置日志级别为 7 存放路径 logs/web.log 按天存放,最多存放 15 天
	beego.SetLogger("file", `{
		"level":7,	
		"filename":"log/zpaas.log",
		"daily":true,
		"maxdays":15
		}`)
}
