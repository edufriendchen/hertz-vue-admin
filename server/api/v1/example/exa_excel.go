package example

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/example"
	"go.uber.org/zap"
	"os"
	"reflect"
	"strings"
)

type ExcelApi struct{}

// /excel/importExcel 接口，与upload接口作用类似，只是把文件存到resource/excel目录下，用于导入Excel时存放Excel文件(ExcelImport.xlsx)
// /excel/loadExcel接口，用于读取resource/excel目录下的文件((ExcelImport.xlsx)并加载为[]model.SysBaseMenu类型的示例数据
// /excel/exportExcel 接口，用于读取前端传来的tableData，生成Excel文件并返回
// /excel/downloadTemplate 接口，用于下载resource/excel目录下的 ExcelTemplate.xlsx 文件，作为导入的模板

// ExportExcel
// @Tags      excel
// @Summary   导出Excel
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/octet-stream
// @Param     data  body  example.ExcelInfo  true  "导出Excel文件信息"
// @Success   200
// @Router    /excel/exportExcel [post]
func (e *ExcelApi) ExportExcel(ctx context.Context, c *app.RequestContext) {
	var excelInfo example.ExcelInfo
	err := c.BindAndValidate(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	filePath := global.CONFIG.Excel.Dir + excelInfo.FileName
	// err = excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	// if err != nil {
	// 	global.LOG.Error("转换Excel失败!", zap.Error(err))
	// 	response.FailWithMessage("转换Excel失败", c)
	// 	return
	// }

	fmt.Println("filePath:", filePath)

	// filePathAndName, err := excelService.ExportExcel(titleList, infoList)
	// filePathAndName, err := utils.InnerExportExcel[excelInfo](infoList)
	//
	// if err != nil {
	// 	fmt.Println("ExportExcel 出错")
	// }
	// fmt.Println(filePathAndName)

	c.Header("success", "true")
}

func FindUtil[T any](i []T) (r []string, err error) {
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(i[0])
	titles := []string{"ID"}
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		item := fieldType.Tag.Get("field")
		if item == "" {
			item = "未配置"
		}
		titles = append(titles, item)
	}
	fmt.Println(titles)
	return titles, nil
}

// ImportExcel
// @Tags      excel
// @Summary   导入Excel文件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                           true  "导入Excel文件"
// @Success   200   {object}  response.Response{msg=string}  "导入Excel文件"
// @Router    /excel/importExcel [post]
func (e *ExcelApi) ImportExcel(ctx context.Context, c *app.RequestContext) {
	header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	_ = c.SaveUploadedFile(header, global.CONFIG.Excel.Dir+"ExcelImport.xlsx")
	response.OkWithMessage("导入成功", c)
}

// LoadExcel
// @Tags      excel
// @Summary   加载Excel数据
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=response.PageResult,msg=string}  "加载Excel数据,返回包括列表,总数,页码,每页数量"
// @Router    /excel/loadExcel [get]
func (e *ExcelApi) LoadExcel(ctx context.Context, c *app.RequestContext) {
	menus, err := excelService.ParseExcel2InfoList()
	if err != nil {
		global.LOG.Error("加载数据失败!", zap.Error(err))
		response.FailWithMessage("加载数据失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menus,
		Total:    int64(len(menus)),
		Page:     1,
		PageSize: 999,
	}, "加载数据成功", c)
}

// DownloadTemplate
// @Tags      excel
// @Summary   下载模板
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     fileName  query  string  true  "模板名称"
// @Success   200
// @Router    /excel/downloadTemplate [get]
func (e *ExcelApi) DownloadTemplate(ctx context.Context, c *app.RequestContext) {
	fileName := c.Query("fileName")
	filePath := global.CONFIG.Excel.Dir + fileName

	fi, err := os.Stat(filePath)
	if err != nil {
		global.LOG.Error("文件不存在!", zap.Error(err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	if fi.IsDir() {
		global.LOG.Error("不支持下载文件夹!", zap.Error(err))
		response.FailWithMessage("不支持下载文件夹", c)
		return
	}
	c.Header("success", "true")
	c.File(filePath)
}
