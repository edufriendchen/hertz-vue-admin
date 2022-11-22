package example

import (
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
)

type ExaCustomer struct {
	global.MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:客户名" vd:"len($)>0"`             // 客户名
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号" vd:"len($)>0"` // 客户手机号
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                                // 管理ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`            // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`                                    // 管理详情
}
