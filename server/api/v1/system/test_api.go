package system

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common/response"
	"go.uber.org/zap"
)

type TestApi struct{}

func (testApi TestApi) GetApiGroupName() string {
	return "test"
}

type test struct {
	Name string `vd:"len($)>1"`
	Age  int    `vd:"$>0"`
}

// Test GetAllOnlineUser 获取在线用户的列表
func (testApi *TestApi) AddTest(ctx context.Context, c *app.RequestContext) {
	fmt.Println("新增测试接口--")
	var l test
	err := c.BindAndValidate(&l)
	fmt.Println(err)
	fmt.Println("name", l.Name)
	fmt.Println("age", l.Age)
	if err != nil {
		fmt.Println("绑定失败")
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err != nil {
		fmt.Println("验证失败")
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = middleware.AuthAndRecord(ctx, c, common.ADD, testApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("新增成功", c)
}

// DeleteTest Test GetAllOnlineUser 获取在线用户的列表
func (testApi *TestApi) DeleteTest(ctx context.Context, c *app.RequestContext) {
	fmt.Println("删除测试接口--")
	err := middleware.AuthAndRecord(ctx, c, common.DELETE, testApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateTest Test GetAllOnlineUser 获取在线用户的列表
func (testApi *TestApi) UpdateTest(ctx context.Context, c *app.RequestContext) {
	fmt.Println("修改测试接口--")
	err := middleware.AuthAndRecord(ctx, c, common.UPDATE, testApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// Test GetAllOnlineUser 获取在线用户的列表
func (testApi *TestApi) GetTest(ctx context.Context, c *app.RequestContext) {
	fmt.Println("获取测试接口--")
	err := middleware.AuthAndRecord(ctx, c, common.GET, testApi.GetApiGroupName())
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("获取成功", c)
}
