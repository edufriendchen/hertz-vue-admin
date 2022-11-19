package system

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	systemReq "github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"go.uber.org/zap"
)

type OnlineUserApi struct{}

// GetAllOnlineUser 获取在线用户的列表
func (onlineUserApi *OnlineUserApi) GetAllOnlineUser(ctx context.Context, c *app.RequestContext) {
	var pageInfo systemReq.SearchApiParams
	err := c.BindAndValidate(&pageInfo)
	fmt.Println(c.GetRequest().Body())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	onlineUserList, total, err := onlineUserService.GetAll()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	onlineUserList = paging(onlineUserList, pageInfo)
	response.OkWithDetailed(response.PageResult{
		List:     onlineUserList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 强制在线用户下线
func (onlineUserApi *OnlineUserApi) DeleteOnlineUser(ctx context.Context, c *app.RequestContext) {
	var delOnlineUserList system.DelOnlineUserList
	err := c.BindAndValidate(&delOnlineUserList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = utils.GetToken(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, valueToken := range delOnlineUserList.TokenList {
		err = onlineUserService.KickOut(valueToken)
		if err != nil {
			global.LOG.Error("强制下线用户失败，目标token:" + valueToken)
		}
	}
	response.Ok(c)
}

// paging 数组分页
func paging(T []system.SysOnlineUserInfo, info systemReq.SearchApiParams) []system.SysOnlineUserInfo {
	limit := info.PageSize
	begin := info.PageSize * (info.Page - 1)
	end := begin + limit
	if begin > len(T) {
		begin = 0
	}
	if end > len(T) {
		end = len(T)
	}
	return T[begin:end]
}

// ExcelOnlineUser onlineUserApi Excel导出在线用户的列表
func (onlineUserApi *OnlineUserApi) ExcelOnlineUser(ctx context.Context, c *app.RequestContext) {
	fmt.Println(c.GetRequest().Body())
	onlineUserList, total, err := onlineUserService.GetAll()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	filePath, err := utils.ExportExcel[system.SysOnlineUserInfo](onlineUserList)
	if err != nil {
		global.LOG.Error("导出失败!", zap.Error(err))
		response.FailWithMessage("导出失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:  onlineUserList,
		Total: total,
	}, "获取成功", c)

	c.Header("success", "true")
	c.File(filePath)
}
