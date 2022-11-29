package middleware

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"strings"
	"time"
)

// LogLayout 日志layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // 存储自定义原数据
	Path      string                 // 访问路径
	Query     string                 // 携带query
	Body      string                 // 携带body数据
	IP        string                 // ip地址
	UserAgent string                 // 代理
	Error     string                 // 错误
	Cost      time.Duration          // 花费时间
	Source    string                 // 来源
}

type Logger struct {
	// Filter 用户自定义过滤
	Filter func(ctx context.Context, c *app.RequestContext) bool
	// FilterKeyword 关键字过滤(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess 鉴权处理
	AuthProcess func(c *app.RequestContext, layout *LogLayout)
	// 日志处理
	Print func(LogLayout)
	// Source 服务唯一标识
	Source string
}

func (l Logger) SetLoggerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		path := c.Request.URI().Path()
		query := c.GetRawData()
		var body []byte
		if l.Filter != nil && !l.Filter(ctx, c) {
			body, _ = c.Body()
			// 将原body塞回去
			c.SetBodyStream(io.NopCloser(bytes.NewBuffer(body)), 1024)
		}
		c.Next(ctx)
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      string(path),
			Query:     string(query),
			IP:        c.ClientIP(),
			UserAgent: string(c.UserAgent()),
			Error:     strings.TrimRight(c.Errors.ByType(ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(ctx, c) {
			layout.Body = string(body)
		}
		if l.AuthProcess != nil {
			// 处理鉴权需要的信息
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// 自行判断key/value 脱敏等
			l.FilterKeyword(&layout)
		}
		// 自行处理日志
		l.Print(layout)
	}
}

func DefaultLogger() app.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			// 标准输出,k8s做收集
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "Hertz-Vue-Admin",
	}.SetLoggerMiddleware()
}
