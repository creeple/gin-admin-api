package main

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/core"
	"gin-admin-api/global"
	"gin-admin-api/router"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println(config.Config.System)
	fmt.Println(config.Config.Logger)

	global.Log = core.InitLogger()
	global.Log.Warnln("日志")
	global.Log.Error("日志")
	global.Log.Infof("日志")

	err := core.MysqlInit()
	if err != nil {
		global.Log.Error(err.Error())
	}
	fmt.Println(config.Config.Mysql)
	_ = router.InitRouter().Run()
}
