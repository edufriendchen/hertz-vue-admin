package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/url"
	"os"
	"strings"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

var caser = cases.Title(language.English)

// PreviewTemp
// @Tags      AutoCode
// @Summary   预览创建后的代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.AutoCodeStruct                                      true  "预览创建代码"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "预览创建后的代码"
// @Router    /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(ctx context.Context, c *app.RequestContext) {
	var a system.AutoCodeStruct
	_ = c.BindAndValidate(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment() // 处理go关键字
	a.PackageT = caser.String(a.Package)
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.LOG.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// CreateTemp
// @Tags      AutoCode
// @Summary   自动代码模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.AutoCodeStruct  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(ctx context.Context, c *app.RequestContext) {
	var a system.AutoCodeStruct
	_ = c.BindAndValidate(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment()
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Error(err))
			c.Header("success", "false")
			c.Header("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	a.PackageT = caser.String(a.Package)
	err := autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, system.ErrAutoMove) {
			c.Header("success", "true")
			c.Header("msg", url.QueryEscape(err.Error()))
		} else {
			c.Header("success", "false")
			c.Header("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Header("Content-Type", "application/json")
		c.Header("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// GetDB
// @Tags      AutoCode
// @Summary   获取当前所有数据库
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前所有数据库"
// @Router    /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(ctx context.Context, c *app.RequestContext) {
	businessDB := c.Query("businessDB")
	dbs, err := autoCodeService.Database(businessDB).GetDB(businessDB)
	var dbList []map[string]interface{}
	for _, db := range global.CONFIG.DBList {
		var item = make(map[string]interface{})
		item["aliasName"] = db.AliasName
		item["dbName"] = db.Dbname
		item["disable"] = db.Disable
		item["dbtype"] = db.Type
		dbList = append(dbList, item)
	}
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, "获取成功", c)
	}
}

// GetTables
// @Tags      AutoCode
// @Summary   获取当前数据库所有表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前数据库所有表"
// @Router    /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(ctx context.Context, c *app.RequestContext) {
	dbName := c.DefaultQuery("dbName", global.CONFIG.Mysql.Dbname)
	businessDB := c.Query("businessDB")
	tables, err := autoCodeService.Database(businessDB).GetTables(businessDB, dbName)
	if err != nil {
		global.LOG.Error("查询table失败!", zap.Error(err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// GetColumn
// @Tags      AutoCode
// @Summary   获取当前表所有字段
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前表所有字段"
// @Router    /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(ctx context.Context, c *app.RequestContext) {
	businessDB := c.Query("businessDB")
	dbName := c.DefaultQuery("dbName", global.CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database(businessDB).GetColumn(businessDB, tableName, dbName)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}

// CreatePackage
// @Tags      AutoCode
// @Summary   创建package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(ctx context.Context, c *app.RequestContext) {
	var a system.SysAutoCode
	_ = c.BindAndValidate(&a)
	if err := utils.Verify(a, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := autoCodeService.CreateAutoCode(&a)
	if err != nil {

		global.LOG.Error("创建成功!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// GetPackage
// @Tags      AutoCode
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(ctx context.Context, c *app.RequestContext) {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"pkgs": pkgs}, "获取成功", c)
	}
}

// DelPackage
// @Tags      AutoCode
// @Summary   删除package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "删除package成功"
// @Router    /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(ctx context.Context, c *app.RequestContext) {
	var a system.SysAutoCode
	_ = c.BindAndValidate(&a)
	err := autoCodeService.DelPackage(a)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// AutoPlug
// @Tags      AutoCode
// @Summary   创建插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建插件模板"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建插件模板成功"
// @Router    /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) AutoPlug(ctx context.Context, c *app.RequestContext) {
	var a system.AutoPlugReq
	err := c.BindAndValidate(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Snake = strings.ToLower(a.PlugName)
	a.NeedModel = a.HasRequest || a.HasResponse
	err = autoCodeService.CreatePlug(a)
	if err != nil {
		global.LOG.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
		return
	}
	response.Ok(c)
}

// InstallPlugin
// @Tags      AutoCode
// @Summary   安装插件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     plug  formData  file                                              true  "this is a test file"
// @Success   200   {object}  response.Response{data=[]interface{},msg=string}  "安装插件成功"
// @Router    /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) InstallPlugin(ctx context.Context, c *app.RequestContext) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := autoCodeService.InstallPlugin(header)
	webStr := "web插件安装成功"
	serverStr := "server插件安装成功"
	if web == -1 {
		webStr = "web端插件未成功安装，请按照文档自行解压安装，如果为纯后端插件请忽略此条提示"
	}
	if server == -1 {
		serverStr = "server端插件未成功安装，请按照文档自行解压安装，如果为纯前端插件请忽略此条提示"
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData([]interface{}{
		gin.H{
			"code": web,
			"msg":  webStr,
		},
		gin.H{
			"code": server,
			"msg":  serverStr,
		}}, c)
}
