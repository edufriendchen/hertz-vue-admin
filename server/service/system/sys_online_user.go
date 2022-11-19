package system

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"time"
)

type OnlineUserService struct{}

// SaveOnlineUserInfo 缓存用户在线状态
func (onlineUserService *OnlineUserService) SaveOnlineUserInfo(c *app.RequestContext, user system.SysUser, token string) error {
	browser := string(c.GetHeader("User-Agent"))
	ip := c.ClientIP()
	arr, err := utils.GetAddressByIp(ip)
	if err != nil {
		arr = "未知"
	}
	loginTime := time.Now()
	sysOnlineUserInfo := system.SysOnlineUserInfo{
		Username:  user.Username,
		NickName:  user.NickName,
		Browser:   browser,
		Ip:        ip,
		Address:   arr,
		LoginTime: loginTime,
		Token:     token,
	}
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	sysOnlineUserInfoJson, err := json.Marshal(sysOnlineUserInfo)
	if err != nil {
		return err
	}
	err = global.REDIS.Set(context.Background(), global.CONFIG.OnlineUser.Key+token, sysOnlineUserInfoJson, timer).Err()
	return nil
}

func (onlineUserService *OnlineUserService) ExistToken(token string) error {
	jsonOnlineUserInfo, err := global.REDIS.Get(context.Background(), global.CONFIG.OnlineUser.Key+token).Result()
	if err != nil || jsonOnlineUserInfo == "" {
		return err
	}
	var onlineUserInfo system.SysOnlineUserInfo
	err = json.Unmarshal([]byte(jsonOnlineUserInfo), &onlineUserInfo)
	if err != nil || len(onlineUserInfo.Token) == 0 {
		return err
	}
	return nil
}

// KickOut 强制在线用户下线
func (onlineUserService *OnlineUserService) KickOut(token string) error {
	// 从redis删去在线状态
	rows, fistErr := global.REDIS.Del(context.Background(), global.CONFIG.OnlineUser.Key+token).Result()
	// 将token拉黑
	jwt := system.JwtBlacklist{Jwt: token}
	err := global.DB.Create(&jwt).Error
	if err != nil && rows == 0 {
		return err
	}
	global.BlackCache.SetDefault(jwt.Jwt, struct{}{})
	if fistErr != nil {
		return err
	}
	return nil
}

// GetAll 获取所有在线用户
func (onlineUserService *OnlineUserService) GetAll() (out []system.SysOnlineUserInfo, total int64, err error) {
	var onlineUserList []system.SysOnlineUserInfo
	test := global.CONFIG.OnlineUser.Key + "*"
	keyList, err := global.REDIS.Keys(context.Background(), test).Result()
	num := len(keyList)
	if err != nil || num == 0 {
		return onlineUserList, total, err
	}
	var item system.SysOnlineUserInfo
	for _, value := range keyList {
		i, _ := global.REDIS.Get(context.Background(), value).Result()
		err := json.Unmarshal([]byte(i), &item)
		if err == nil {
			onlineUserList = append(onlineUserList, item)
		}
	}
	return onlineUserList, int64(num), nil
}

/**
 * 检测用户是否在之前是否已经登录
 * @param userName 用户名
 */
func (onlineUserService *OnlineUserService) CheckLoginOnUser(userName string) (total int, err error) {
	onlineUserList, _, err := onlineUserService.GetAll()
	if err != nil {
		return 0, nil
	}
	num := 0
	jwtBlacklist := system.JwtBlacklist{Jwt: ""}
	for _, value := range onlineUserList {
		if value.Username == userName {
			num++
			// 1、从Redis中删除在线信息
			_, itemErr := global.REDIS.Del(context.Background(), global.CONFIG.OnlineUser.Key+value.Token).Result()
			if itemErr != nil {
				global.LOG.Error("从Redis中删除用户:" + userName + "旧token：" + value.Token + "错误！")
			}
			// 2、 将旧登录token放入黑名单
			jwtBlacklist.Jwt = value.Token
			err = global.DB.Create(&jwtBlacklist).Error
			if err != nil {
				return
			}
		}
	}
	return num, nil
}
