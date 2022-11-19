package system

type SysRolesApiGroup struct {
	ApiGroupId uint   `json:"apiGroupId" gorm:"index;comment:菜单ID"`
	RoleId     uint   `json:"RoleId" gorm:"index;comment:角色ID"`
	Permission []byte `gorm:"comment:权限组"`
}

func (rolesApiGroup *SysRolesApiGroup) TableName() string {
	return "sys_roles_api_group"
}
