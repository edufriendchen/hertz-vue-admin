package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	systemRes "github.com/edufriendchen/hertz-vue-admin/server/model/system/response"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// UpdateCasbin
// @Tags      Casbin
// @Summary   更新角色api权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CasbinInReceive        true  "权限id, 权限模型列表"
// @Success   200   {object}  response.Response{msg=string}  "更新角色api权限"
// @Router    /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(ctx context.Context, c *app.RequestContext) {
	var cmr request.CasbinInReceive
	err := c.BindAndValidate(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetPolicyPathByAuthorityId
// @Tags      Casbin
// @Summary   获取权限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CasbinInReceive                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  response.Response{data=systemRes.PolicyPathResponse,msg=string}  "获取权限列表,返回包括casbin详情列表"
// @Router    /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(ctx context.Context, c *app.RequestContext) {
	var casbin request.CasbinInReceive
	err := c.BindAndValidate(&casbin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
