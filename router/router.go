/**
 * @Author: lixiang
 * @Date: 2025/3/11 20:38
 * @Description: router.go
 */

package router

import (
	"gin-admin-api/api"
	"gin-admin-api/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	router.Use(gin.Recovery())

	register(router)
	return router
}

func register(router *gin.Engine) {
	//todo 接口url

	router.GET("/api/testSuccess", api.TestSuccess)
	router.GET("/api/testFail", api.TestFail)
}
