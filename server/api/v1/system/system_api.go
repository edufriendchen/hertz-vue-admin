package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	systemRes "github.com/edufriendchen/hertz-vue-admin/server/model/system/response"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// GetSystemConfig
// @Tags      System
// @Summary   获取配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysConfigResponse,msg=string}  "获取配置文件内容,返回包括系统配置"
// @Router    /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(ctx context.Context, c *app.RequestContext) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
}

// SetSystemConfig
// @Tags      System
// @Summary   设置配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      system.System                   true  "设置配置文件内容"
// @Success   200   {object}  response.Response{data=string}  "设置配置文件内容"
// @Router    /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(ctx context.Context, c *app.RequestContext) {
	var sys system.System
	err := c.BindAndValidate(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = systemConfigService.SetSystemConfig(sys)
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// ReloadSystem
// @Tags      System
// @Summary   重启系统
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "重启系统"
// @Router    /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(ctx context.Context, c *app.RequestContext) {
	err := utils.Reload()
	if err != nil {
		global.LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
		return
	}
	response.OkWithMessage("重启系统成功", c)
}

// GetServerInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(ctx context.Context, c *app.RequestContext) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}
