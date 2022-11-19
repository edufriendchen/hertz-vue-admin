package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *server.Hertz) (R *route.RouterGroup) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
