package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
	"github.com/edufriendchen/hertz-vue-admin/server/plugin/email/api"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *route.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord)
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
