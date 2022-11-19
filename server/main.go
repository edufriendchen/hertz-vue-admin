package main

import (
	"go.uber.org/zap"

	"github.com/edufriendchen/hertz-vue-admin/server/core"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath
//
// /
func main() {
	global.VIPER = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.DB != nil {
		initialize.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
