package email

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/hertz-vue-admin/server/plugin/email/global"
	"github.com/edufriendchen/hertz-vue-admin/server/plugin/email/router"
)

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Nickname = Nickname
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	return &emailPlugin{}
}

func (*emailPlugin) Register(group *route.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
