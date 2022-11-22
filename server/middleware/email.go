package middleware

import "C"
import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"strconv"
	"time"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/service"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	utils2 "github.com/edufriendchen/hertz-vue-admin/server/utils"
	"go.uber.org/zap"
)

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func ErrorToEmail() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var username string
		claims, _ := utils2.GetClaims(c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			user, err := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body, _ := io.ReadAll(c.GetRequest().BodyStream())
		// 再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据

		c.Request.SetBody(body)

		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: string(c.GetRequest().Method()),
			Path:   string(c.GetRequest().URI().Path()),
			Agent:  string(c.UserAgent()),
			Body:   string(body),
		}
		now := time.Now()

		c.Next(ctx)

		latency := time.Since(now)
		status := c.GetResponse().StatusCode()
		record.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
