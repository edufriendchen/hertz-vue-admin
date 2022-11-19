package system

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/request"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	systemReq "github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	systemRes "github.com/edufriendchen/hertz-vue-admin/server/model/system/response"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

func (authorityApi AuthorityApi) GetApiGroupName() string {
	return "role"
}

// CreateAuthority
// @Tags      Authority
// @Summary   创建角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
// @Router    /authority/createAuthority [post]
func (authorityApi *AuthorityApi) CreateAuthority(ctx context.Context, c *app.RequestContext) {
	err := middleware.AuthAndRecord(ctx, c, common.ADD, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authority system.SysAuthority
	err = c.BindAndValidate(&authority)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(authority, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// CopyAuthority
// @Tags      Authority
// @Summary   拷贝角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      response.SysAuthorityCopyResponse                                  true  "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "拷贝角色,返回包括系统角色详情"
// @Router    /authority/copyAuthority [post]
func (authorityApi *AuthorityApi) CopyAuthority(ctx context.Context, c *app.RequestContext) {
	err := middleware.AuthAndRecord(ctx, c, common.ADD, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var copyInfo systemRes.SysAuthorityCopyResponse
	err = c.BindAndValidate(&copyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo, utils.OldAuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo.Authority, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authBack, err := authorityService.CopyAuthority(copyInfo)
	if err != nil {
		global.LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
}

// DeleteAuthority
// @Tags      Authority
// @Summary   删除角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority            true  "删除角色"
// @Success   200   {object}  response.Response{msg=string}  "删除角色"
// @Router    /authority/deleteAuthority [post]
func (authorityApi *AuthorityApi) DeleteAuthority(ctx context.Context, c *app.RequestContext) {
	err := middleware.AuthAndRecord(ctx, c, common.UPDATE, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authority system.SysAuthority
	err = c.BindAndValidate(&authority)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(authority, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityService.DeleteAuthority(&authority)
	if err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateAuthority
// @Tags      Authority
// @Summary   更新角色信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "更新角色信息,返回包括系统角色详情"
// @Router    /authority/updateAuthority [post]
func (authorityApi *AuthorityApi) UpdateAuthority(ctx context.Context, c *app.RequestContext) {
	err := middleware.AuthAndRecord(ctx, c, common.UPDATE, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var auth system.SysAuthority
	err = c.BindAndValidate(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authority, err := authorityService.UpdateAuthority(auth)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
}

// GetAuthorityList
// @Tags      Authority
// @Summary   分页获取角色列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router    /authority/getAuthorityList [post]
func (authorityApi *AuthorityApi) GetAuthorityList(ctx context.Context, c *app.RequestContext) {

	fmt.Println("Content-Type", string(c.GetHeader("Content-Type")))
	fmt.Println("Content-Type", string(c.GetRequest().Body()))

	err := middleware.AuthAndRecord(ctx, c, common.GET, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var pageInfo request.PageInfo
	err = c.BindAndValidate(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := authorityService.GetAuthorityInfoList(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SetDataAuthority
// @Tags      Authority
// @Summary   设置角色资源权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority            true  "设置角色资源权限"
// @Success   200   {object}  response.Response{msg=string}  "设置角色资源权限"
// @Router    /authority/setDataAuthority [post]
func (authorityApi *AuthorityApi) SetDataAuthority(ctx context.Context, c *app.RequestContext) {
	err := middleware.AuthAndRecord(ctx, c, common.ADD, authorityApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var auth system.SysAuthority
	err = c.BindAndValidate(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityService.SetDataAuthority(auth)
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
