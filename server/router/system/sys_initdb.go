package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *server.Hertz) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkdb", dbApi.CheckDB) // 检测是否需要初始化数据库
	}
}
