package system

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type TestRouter struct{}

func (o *OnlineUserRouter) InitTestRouter(Router *route.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.SystemApiGroup.TestApi
	{
		testRouter.POST("addTest", testApi.AddTest)         // 获取所有在线用户
		testRouter.GET("getTest", testApi.GetTest)          // 获取所有在线用户
		testRouter.DELETE("deleteTest", testApi.DeleteTest) // 获取所有在线用户
		testRouter.PUT("updateTest", testApi.UpdateTest)    // 获取所有在线用户
	}
}
