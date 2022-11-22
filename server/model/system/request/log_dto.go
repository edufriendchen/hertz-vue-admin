package request

import "github.com/edufriendchen/hertz-vue-admin/server/model/common/request"

type LogDto struct {
	Ip     string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Method string `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Path   string `json:"path" form:"path" gorm:"column:path;comment:请求路径"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:请求状态"`
	UserID int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	request.PageInfo
}
