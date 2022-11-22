package example

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/request"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/example"
	exampleRes "github.com/edufriendchen/hertz-vue-admin/server/model/example/response"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"go.uber.org/zap"
)

type CustomerApi struct{}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(ctx context.Context, c *app.RequestContext) {
	var customer example.ExaCustomer
	err := c.BindAndValidate(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = customerService.CreateExaCustomer(customer)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *CustomerApi) DeleteExaCustomer(ctx context.Context, c *app.RequestContext) {
	var customer example.ExaCustomer
	err := c.BindAndValidate(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.DeleteExaCustomer(customer)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *CustomerApi) UpdateExaCustomer(ctx context.Context, c *app.RequestContext) {
	var customer example.ExaCustomer
	err := c.BindAndValidate(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.UpdateExaCustomer(&customer)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
func (e *CustomerApi) GetExaCustomer(ctx context.Context, c *app.RequestContext) {
	var customer example.ExaCustomer
	err := c.BindAndValidate(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaCustomerResponse{Customer: data}, "获取成功", c)
}

// GetExaCustomerList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *CustomerApi) GetExaCustomerList(ctx context.Context, c *app.RequestContext) {
	var pageInfo request.PageInfo
	err := c.BindAndValidate(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
