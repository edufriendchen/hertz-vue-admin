package system

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
)

type JwtService struct{}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: JsonInBlacklist
// @description: 拉黑jwt
// @param: jwtList model.JwtBlacklist
// @return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: IsBlacklist
// @description: 判断JWT是否在黑名单内部
// @param: jwt string
// @return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetRedisJWT
// @description: 从redis取jwt
// @param: userName string
// @return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(token string) (redisToken string, err error) {
	jsonOnlineUserInfo, err := global.REDIS.Get(context.Background(), global.CONFIG.OnlineUser.Key+token).Result()
	if err == nil {
		return "", err
	}
	var onlineUserInfo system.SysOnlineUserInfo
	err = json.Unmarshal([]byte(jsonOnlineUserInfo), &onlineUserInfo)
	if err != nil || len(onlineUserInfo.Token) == 0 {
		return "", err
	}
	return onlineUserInfo.Token, nil
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, token string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), global.CONFIG.OnlineUser.Key+token, jwt, timer).Err()
	return err
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error

func LoadAll() {
	var data []string
	err := global.DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
