package system

import (
	"github.com/edufriendchen/hertz-vue-admin/server/global"
)

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
