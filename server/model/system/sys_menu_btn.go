package system

import "github.com/edufriendchen/hertz-vue-admin/server/global"

type SysBaseMenuBtn struct {
	global.MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
