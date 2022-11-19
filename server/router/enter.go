package router

import (
	"github.com/edufriendchen/hertz-vue-admin/server/router/example"
	"github.com/edufriendchen/hertz-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
