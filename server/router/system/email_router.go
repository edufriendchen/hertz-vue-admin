package system

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *route.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord)
	EmailApi := v1.ApiGroupApp.SystemApiGroup.EmailApi.EmailTest
	SendEmail := v1.ApiGroupApp.SystemApiGroup.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
