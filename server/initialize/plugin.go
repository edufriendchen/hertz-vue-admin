package initialize

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
	"github.com/edufriendchen/hertz-vue-admin/server/plugin/email"
	"github.com/edufriendchen/hertz-vue-admin/server/utils/plugin"
)

func PluginInit(group *route.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(h *server.Hertz) {
	PublicGroup := h.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := h.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PublicGroup, email.CreateEmailPlug(
		global.CONFIG.Email.To,
		global.CONFIG.Email.From,
		global.CONFIG.Email.Host,
		global.CONFIG.Email.Secret,
		global.CONFIG.Email.Nickname,
		global.CONFIG.Email.Port,
		global.CONFIG.Email.IsSSL,
	))
}
