package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type LogRouter struct{}

func (o *OnlineUserRouter) InitLogRouter(Router *server.Hertz) route.IRoutes {
	onlineUserRouterWithoutRecord := Router.Group("logs")
	logApi := v1.ApiGroupApp.SystemApiGroup.LogApi
	{
		onlineUserRouterWithoutRecord.POST("", logApi.QueryLog)
		onlineUserRouterWithoutRecord.POST("error", logApi.QueryErrorLog)
		onlineUserRouterWithoutRecord.DELETE("del/info", logApi.DelAllInfoLog)
		onlineUserRouterWithoutRecord.DELETE("del/error", logApi.DelAllErrorLog)
	}
	return onlineUserRouterWithoutRecord
}
