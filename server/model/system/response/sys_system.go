package response

import "github.com/edufriendchen/hertz-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
