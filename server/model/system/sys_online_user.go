package system

import (
	"time"
)

type SysOnlineUserInfo struct {
	Username  string    `json:"userName" excel:"用户名"`
	NickName  string    `json:"nickName" excel:"用户昵称"`
	Browser   string    `json:"browser" excel:"浏览器"`
	Ip        string    `json:"ip" excel:"登录IP"`
	Address   string    `json:"address" excel:"登录地点"`
	LoginTime time.Time `json:"login" excel:"时间"`
	Token     string    `json:"token" excel:"凭证"`
}

type DelOnlineUserList struct {
	TokenList []string `json:"tokenList"`
}

func (s SysOnlineUserInfo) GetValue() []string {
	return []string{
		s.Username,
		s.NickName,
		s.Browser,
		s.Ip,
		s.Address,
		s.LoginTime.String(),
		s.Token,
	}
}
