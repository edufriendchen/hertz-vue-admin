package utils

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	systemReq "github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	uuid "github.com/satori/go.uuid"
)

func IsToken(token string) (string, error) {
	j := NewJWT()
	_, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return "", err
	}
	return token, err
}

func GetToken(c *app.RequestContext) (string, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	_, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return "", err
	}
	return token, err
}

func GetClaims(c *app.RequestContext) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *app.RequestContext) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *app.RequestContext) (uuid.UUID, error) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			fmt.Println("1:")
			return uuid.UUID{}, err
		} else {
			fmt.Println("2:", cl.UUID)
			return cl.UUID, nil
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		fmt.Println("3:", waitUse.UUID)
		return waitUse.UUID, nil
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuidByToken(token string) (string, error) {
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return "", err
	}
	return claims.UUID.String(), nil
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *app.RequestContext) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *app.RequestContext) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// ----------------------------------------------------
func GetClaims2(c *app.RequestContext) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
