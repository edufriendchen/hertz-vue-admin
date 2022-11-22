package system

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/request"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	systemReq "github.com/edufriendchen/hertz-vue-admin/server/model/system/request"
	systemRes "github.com/edufriendchen/hertz-vue-admin/server/model/system/response"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(ctx context.Context, c *app.RequestContext) {
	var l systemReq.Login
	err := c.BindAndValidate(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(ctx, c, *user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(ctx context.Context, c *app.RequestContext, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 多点登录
	if !global.CONFIG.System.SinglePointLogin {
		global.LOG.Info("[多点登录]-可在配置文件中切换。", zap.Error(err))
		// 缓存在线用户信息
		err = onlineUserService.SaveOnlineUserInfo(c, user, token)
		if err != nil {
			global.LOG.Error("缓存在线用户信息失败!", zap.Error(err))
		}

		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	// 多点登录  限制登录数量版本
	if !global.CONFIG.System.SinglePointLogin {
		global.LOG.Info("[多点登录]-可在配置文件中切换。", zap.Error(err))
		// 查看该账号是否已有登录
		onlineUserList, _, err := onlineUserService.GetAll()
		num := 0
		for _, value := range onlineUserList {
			if value.Username == user.Username {
				num++
			}
		}
		// 该账号存在登录
		if err != redis.Nil && num != 0 {
			fmt.Println("该账号现有登录量", err, num, global.CONFIG.System.MaxMultipointLoginNum)
			if err == nil && num > global.CONFIG.System.MaxMultipointLoginNum-1 {
				global.LOG.Error("超过账号最大登录限制!", zap.Error(err))
				response.FailWithMessage("超过账号最大登录限制!", c)
				return
			}
			// 缓存在线用户信息
			err = onlineUserService.SaveOnlineUserInfo(c, user, token)
			if err != nil {
				global.LOG.Error("缓存在线用户信息失败!", zap.Error(err))
			}

			response.OkWithDetailed(systemRes.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			}, "登录成功", c)
			return
		}

		// 缓存在线用户信息
		err = onlineUserService.SaveOnlineUserInfo(c, user, token)
		if err != nil {
			global.LOG.Error("缓存在线用户信息失败!", zap.Error(err))
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	// 单点登录
	global.LOG.Info("[单点登录模式]-可在配置文件中切换。", zap.Error(err))
	{
		onlineUserList, _, err := onlineUserService.GetAll()
		// 如果在线信息为空，则直接Redis缓存用户在线信息，用户登录成功
		if err == redis.Nil {
			// 单点登录状态下，用户第一次登录，缓存在线信息
			err = onlineUserService.SaveOnlineUserInfo(c, user, token)
			if err != nil {
				global.LOG.Error("单点登录状态，缓存在线用户信息失败!", zap.Error(err))
				response.FailWithMessage("设置登录状态失败", c)
				return
			}
			response.OkWithDetailed(systemRes.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			}, "登录成功", c)
		}

		// 如果已经存在在线信息
		num := 0
		for _, value := range onlineUserList {
			if value.Username == user.Username {
				num++
				// 1、从Redis中删除在线信息
				delStatus, itemErr := global.REDIS.Del(context.Background(), global.CONFIG.OnlineUser.Key+value.Token).Result()
				fmt.Println("删除的用户信息：", delStatus)
				if itemErr != nil {
					global.LOG.Error("从Redis中删除用户:" + user.Username + "旧token：" + value.Token + "错误！")
				}
				// 2、 将旧登录token放入黑名单
			}
		}
		fmt.Println("该用户已经存在的在线信息数量", num)
		err = onlineUserService.SaveOnlineUserInfo(c, user, token)
		if err != nil {
			global.LOG.Error("单点登录状态，缓存在线用户信息失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)

	}

	// 单点登录
	// global.LOG.Info("[单点登录模式]-可在配置文件中切换。", zap.Error(err))
	// if jwtStr, err := jwtService.GetRedisJWT(token); err == redis.Nil {
	// 	fmt.Println("action-1")
	// 	// 单点登录状态下，用户第一次登录，缓存在线信息
	// 	err = onlineUserService.SaveOnlineUserInfo(c, user, token)
	// 	if err != nil {
	// 		global.LOG.Error("单点登录状态，缓存在线用户信息失败!", zap.Error(err))
	// 		response.FailWithMessage("设置登录状态失败", c)
	// 		return
	// 	}
	// 	response.OkWithDetailed(systemRes.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// } else if err != nil {
	// 	fmt.Println("action-2")
	//
	// 	global.LOG.Error("单点登录状态，查询Redis是否存在属于该用户的token过程出错", zap.Error(err))
	// 	response.FailWithMessage("设置登录状态失败", c)
	// } else {
	// 	fmt.Println("action-3")
	// 	global.LOG.Error("单点登录状态，设置登录状态失败", zap.Error(err))
	// 	var blackJWT system.JwtBlacklist
	// 	blackJWT.Jwt = jwtStr
	// 	// 将旧token添加进黑名单
	// 	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	// 		response.FailWithMessage("jwt作废失败", c)
	// 		return
	// 	}
	// 	err = onlineUserService.SaveOnlineUserInfo(c, user, token)
	// 	if err != nil {
	// 		global.LOG.Error("缓存在线用户信息失败!", zap.Error(err))
	// 	}
	//
	// 	response.OkWithDetailed(systemRes.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// }
}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(ctx context.Context, c *app.RequestContext) {
	var r systemReq.Register
	err := c.BindAndValidate(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(ctx context.Context, c *app.RequestContext) {
	var req systemReq.ChangePasswordReq
	err := c.BindAndValidate(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	c.Set("UserId", uid)
	u := &system.SysUser{MODEL: global.MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(ctx context.Context, c *app.RequestContext) {
	var pageInfo request.PageInfo
	err := c.BindAndValidate(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// ExportUserList ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ExportUserList(ctx context.Context, c *app.RequestContext) {
	var user system.SysUser
	err := c.BindAndValidate(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	db := global.DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Preload("Authorities").Preload("Authority").Find(&userList).Error
	filePath, err := utils.ExportExcel[system.SysUser](userList)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	c.Header("success", "true")
	c.File(filePath)
	response.OkWithMessage("获取成功", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(ctx context.Context, c *app.RequestContext) {
	var sua systemReq.SetUserAuth
	err := c.BindAndValidate(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
		response.OkWithMessage("修改成功", c)
	}
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(ctx context.Context, c *app.RequestContext) {
	var sua systemReq.SetUserAuthorities
	err := c.BindAndValidate(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(ctx context.Context, c *app.RequestContext) {
	var reqId request.GetById
	err := c.BindAndValidate(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(ctx context.Context, c *app.RequestContext) {
	var user systemReq.ChangeUserInfo
	err := c.BindAndValidate(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		MODEL: global.MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(ctx context.Context, c *app.RequestContext) {
	var user systemReq.ChangeUserInfo
	err := c.BindAndValidate(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetUserInfo(system.SysUser{
		MODEL: global.MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(ctx context.Context, c *app.RequestContext) {
	uuid, err := utils.GetUserUuid(c)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(ctx context.Context, c *app.RequestContext) {
	var user system.SysUser
	err := c.BindAndValidate(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}
