package system

type SysApiGroup struct {
	ID   uint   `json:"id" excel:"API组ID"`
	Name string `json:"name" excel:"API组名"`
	Path string `json:"path" excel:"API路径"`
}

func (s *SysApiGroup) TableName() string {
	return "sys_api_group"
}
