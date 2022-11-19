package utils

import (
	"fmt"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

func ExportExcel[T system.MyType](list []T) (filePath string, err error) {
	excel := excelize.NewFile()
	typeOfCat := reflect.TypeOf(list[0])
	titles := []string{"ID"}
	var value []string
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		item := fieldType.Tag.Get("excel")
		if item != "" {
			titles = append(titles, item)
		}
	}
	fmt.Println("titles", titles)
	err = excel.SetSheetRow("Sheet1", "A1", &titles)
	if err != nil {
		return "", err
	}
	for index, item := range list {
		value = []string{strconv.Itoa(index + 1)}
		axis := fmt.Sprintf("A%d", index+2)
		value = append(value, item.GetValue()...)
		_ = excel.SetSheetRow("Sheet1", axis, &value)
	}
	filePathAndName := global.CONFIG.Excel.Dir + uuid.NewV4().String() + ".xlsx"
	err = excel.SaveAs(filePathAndName)
	return filePathAndName, err
}
