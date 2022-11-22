package request

import (
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:"用户名" vd:"len($)>1"`
	Password     string `json:"passWord" example:"密码" vd:"len($)>1"`
	NickName     string `json:"nickName" example:"昵称" vd:"len($)>1"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id" vd:"len($)>1"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
}

// User login structure
type Login struct {
	Username  string `json:"username" vd:"len($)>1"`  // 用户名
	Password  string `json:"password" vd:"len($)>1"`  // 密码
	Captcha   string `json:"captcha" vd:"len($)>1"`   // 验证码
	CaptchaId string `json:"captchaId" vd:"len($)>1"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`                         // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password" vd:"len($)>1"`    // 密码
	NewPassword string `json:"newPassword" vd:"len($)>1"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId" vd:"$>0"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey" vd:"$>0"`                                                                  // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           // 冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
