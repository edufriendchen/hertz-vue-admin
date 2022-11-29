package core

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/mattn/go-colorable"
	"time"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/initialize"
	"github.com/edufriendchen/hertz-vue-admin/server/service/system"
	"go.uber.org/zap"
)

func RunWindowsServer() {

	if global.CONFIG.System.SinglePointLogin || global.CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	// 从db加载jwt黑名单数据
	if global.DB != nil {
		system.LoadAll()
	}
	// 初始化日志配置
	ForceConsoleColor()
	DefaultWriter = colorable.NewColorableStdout()
	// 初始化服务配置
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(server.WithExitWaitTime(24 * 60 * 60 * 15))

	s.Use(Logger())
	// 初始化路由
	initialize.Routers(s)

	// Router.Static("/form-generator", "./resource/page")
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Println(
		"| |  | |         | |          \\ \\    / /                 /\\      | |         (_)      \n" +
			"| |__| | ___ _ __| |_ _________\\ \\  / /   _  ___ ______ /  \\   __| |_ __ ___  _ _ __  \n" +
			"|  __  |/ _ \\ '__| __|_  /______\\ \\/ / | | |/ _ \\______/ /\\ \\ / _` | '_ ` _ \\| | '_ \\ \n" +
			"| |  | |  __/ |  | |_ / /        \\  /| |_| |  __/     / ____ \\ (_| | | | | | | | | | |\n" +
			"|_|  |_|\\___|_|   \\__/___|        \\/  \\__,_|\\___|    /_/    \\_\\__,_|_| |_| |_|_|_| |_|\n")
	s.Spin()
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf("Hertz-Vue-Admin", address)
	// global.LOG.Error()
}
