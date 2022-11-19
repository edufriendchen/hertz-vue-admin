package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/service"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if global.CONFIG.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			// 获取请求的PATH
			obj := c.URI().Path()
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next(ctx)
	}
}
