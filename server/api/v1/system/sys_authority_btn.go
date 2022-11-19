package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

// GetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   获取权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAuthorityBtnReq                                      true  "菜单id, 角色id, 选中的按钮id"
// @Success   200   {object}  response.Response{data=response.SysAuthorityBtnRes,msg=string}  "返回列表成功"
// @Router    /authorityBtn/getAuthorityBtn [post]
func (a *AuthorityBtnApi) GetAuthorityBtn(ctx context.Context, c *app.RequestContext) {
	var req request.SysAuthorityBtnReq
	err := c.BindAndValidate(&req)
	res, err := authorityBtnService.GetAuthorityBtn(req)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(res, "查询成功", c)
}

// SetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   设置权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAuthorityBtnReq     true  "菜单id, 角色id, 选中的按钮id"
// @Success   200   {object}  response.Response{msg=string}  "返回列表成功"
// @Router    /authorityBtn/setAuthorityBtn [post]
func (a *AuthorityBtnApi) SetAuthorityBtn(ctx context.Context, c *app.RequestContext) {
	var req request.SysAuthorityBtnReq
	err := c.BindAndValidate(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityBtnService.SetAuthorityBtn(req)
	if err != nil {
		global.LOG.Error("分配失败!", zap.Error(err))
		response.FailWithMessage("分配失败", c)
		return
	}
	response.OkWithMessage("分配成功", c)
}

// CanRemoveAuthorityBtn
// @Tags      AuthorityBtn
// @Summary   设置权限按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "删除成功"
// @Router    /authorityBtn/canRemoveAuthorityBtn [post]
func (a *AuthorityBtnApi) CanRemoveAuthorityBtn(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err := authorityBtnService.CanRemoveAuthorityBtn(id)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
