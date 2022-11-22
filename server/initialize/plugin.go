package initialize

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
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
}
