package system

import (
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/satori/go.uuid"
	"strconv"
)

type SysUser struct {
	global.MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string         `json:"userName" gorm:"index;comment:用户登录名" excel:"用户名"`                                      // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称" excel:"用户昵称"`                               // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色" excel:"用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号" excel:"用户手机号"`                    // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱" excel:"用户邮箱"`                      // 用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结" excel:"用户状态"` // 用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s SysUser) GetValue() []string {
	return []string{
		s.Username,
		s.NickName,
		s.Authority.AuthorityName,
		s.Phone,
		s.Email,
		strconv.Itoa(s.Enable),
	}
}
