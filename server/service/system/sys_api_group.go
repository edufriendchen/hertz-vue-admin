package system

import (
	"errors"
	"fmt"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
)

type SysApiGrop struct{}

func SaveSysApiGrop(apiGrop system.SysApiGroup) error {
	var err error
	err = global.DB.Save(&apiGrop).Error
	return err
}

func GetSysApiGrop(apiGroupId uint) (system.SysApiGroup, error) {
	var apiGroup system.SysApiGroup
	err := global.DB.First(&apiGroup, "id = ?", apiGroupId).Error
	return apiGroup, err
}

func getPermission(roleId uint, apiGroupId uint) (system.SysRolesApiGroup, error) {
	var rolesApiGroup system.SysRolesApiGroup
	err := global.DB.First(&rolesApiGroup, "role_id = ? AND api_group_id = ?", roleId, apiGroupId).Error
	return rolesApiGroup, err
}

func CheckAuth(permissionType common.PermissionType, roleId uint, apiGroupName string) error {
	var apiGroup system.SysApiGroup
	err := global.DB.First(&apiGroup, "name = ?", apiGroupName).Error
	if err != nil {
		return errors.New("不存在该API组")
	}
	fmt.Println("API组名称:", apiGroupName, "对应ID:", apiGroup.ID)
	rolesApiGroup, err := getPermission(roleId, apiGroup.ID)
	if err != nil {
		return err
	}
	fmt.Println("权限对应表", rolesApiGroup.Permission)
	if rolesApiGroup.Permission[permissionType] == 48 {
		return errors.New("没有权限")
	}
	return nil
}
