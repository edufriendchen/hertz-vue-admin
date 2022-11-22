package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *server.Hertz) {
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouterWithoutRecord.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouterWithoutRecord.DELETE("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouterWithoutRecord.PUT("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouterWithoutRecord.POST("copyAuthority", authorityApi.CopyAuthority)       // 拷贝角色
		authorityRouterWithoutRecord.POST("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
