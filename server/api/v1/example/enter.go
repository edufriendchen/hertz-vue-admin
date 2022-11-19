package example

import "github.com/edufriendchen/hertz-vue-admin/server/service"

type ApiGroup struct {
	ExcelApi
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	excelService                 = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
