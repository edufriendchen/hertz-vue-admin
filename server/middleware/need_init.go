package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func NeedInit() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if global.DB == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", c)
			c.Abort()
		} else {
			c.Next(ctx)
		}
		// 处理请求
	}
}
