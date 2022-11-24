package system

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type OnlineUserRouter struct{}

func (o *OnlineUserRouter) InitOnlineUserRouter(Router *route.RouterGroup) route.IRoutes {
	onlineUserRouterWithoutRecord := Router.Group("onlineUser")
	onlineUserApi := v1.ApiGroupApp.SystemApiGroup.OnlineUserApi
	{
		onlineUserRouterWithoutRecord.POST("getAllOnlineUser", onlineUserApi.GetAllOnlineUser) // 获取所有在线用户
		onlineUserRouterWithoutRecord.POST("deleteOnlineUser", onlineUserApi.DeleteOnlineUser) // 强制下线用户失败
		onlineUserRouterWithoutRecord.POST("excelOnlineUser", onlineUserApi.ExcelOnlineUser)
	}
	return onlineUserRouterWithoutRecord
}
