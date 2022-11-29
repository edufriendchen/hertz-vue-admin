package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/request"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	systemReq "github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// CreateSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   创建SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysOperationRecord      true  "创建SysOperationRecord"
// @Success   200   {object}  response.Response{msg=string}  "创建SysOperationRecord"
// @Router    /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(ctx context.Context, c *app.RequestContext) {
	var sysOperationRecord system.SysOperationRecord
	err := c.BindAndValidate(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.CreateSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysOperationRecord      true  "SysOperationRecord模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(ctx context.Context, c *app.RequestContext) {
	var sysOperationRecord system.SysOperationRecord
	err := c.BindAndValidate(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysOperationRecordByIds
// @Tags      SysOperationRecord
// @Summary   批量删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "批量删除SysOperationRecord"
// @Success   200   {object}  response.Response{msg=string}  "批量删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(ctx context.Context, c *app.RequestContext) {
	var IDS request.IdsReq
	err := c.BindAndValidate(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// FindSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   用id查询SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysOperationRecord                                  true  "Id"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysOperationRecord"
// @Router    /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(ctx context.Context, c *app.RequestContext) {
	var sysOperationRecord system.SysOperationRecord
	err := c.BindAndValidate(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(utils.H{"reSysOperationRecord": reSysOperationRecord}, "查询成功", c)
}

// GetSysOperationRecordList
// @Tags      SysOperationRecord
// @Summary   分页获取SysOperationRecord列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysOperationRecordSearch                        true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(ctx context.Context, c *app.RequestContext) {
	var sysOperationRecordSearch systemReq.SysOperationRecordSearch
	err := c.Bind(&sysOperationRecordSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = c.Validate(&sysOperationRecordSearch.PageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := operationRecordService.GetSysOperationRecordInfoList(sysOperationRecordSearch)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     sysOperationRecordSearch.Page,
		PageSize: sysOperationRecordSearch.PageSize,
	}, "获取成功", c)
}
