package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *server.Hertz) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord)
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
	}
}
