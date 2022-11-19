package request

import (
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/request"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
