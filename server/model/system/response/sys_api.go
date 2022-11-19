package response

import "github.com/edufriendchen/hertz-vue-admin/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"controller"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
