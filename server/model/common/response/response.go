package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result2(code int, data interface{}, msg string, ctx *app.RequestContext) {
	// 开始时间
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *app.RequestContext) {
	Result2(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}

func OkWithMessage(message string, ctx *app.RequestContext) {
	Result2(SUCCESS, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx *app.RequestContext) {
	Result2(SUCCESS, data, "查询成功", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *app.RequestContext) {
	Result2(SUCCESS, data, message, ctx)
}

func Fail(ctx *app.RequestContext) {
	Result2(ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func FailWithMessage(message string, ctx *app.RequestContext) {
	Result2(ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx *app.RequestContext) {
	Result2(ERROR, data, message, ctx)
}
