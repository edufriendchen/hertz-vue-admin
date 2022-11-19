package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/edufriendchen/hertz-vue-admin/server/model/common"
	system2 "github.com/edufriendchen/hertz-vue-admin/server/service/system"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system"
	"github.com/edufriendchen/hertz-vue-admin/server/service"
	"github.com/edufriendchen/hertz-vue-admin/server/utils"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

type responseBodyWriter struct {
	ResponseWriter
	body *bytes.Buffer
}
type ResponseWriter interface {
	Malloc(n int) (buf []byte, err error)
	WriteBinary(b []byte) (n int, err error)
	Flush() error
}

func (r responseBodyWriter) Write2(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.WriteBinary(b)
}

const ErrorTypePrivate errors.ErrorType = 1 << 0

func OperationRecord(ctx context.Context, c *app.RequestContext) {
	fmt.Println("OperationRecordTest")

	fmt.Println("c.Request.URI().Path():", c.Request.URI().Path())
	var body []byte
	var userId int
	if string(c.Request.Method()) != http.MethodGet {
		fmt.Println("POST")
		var err error
		body, err = c.Body()
		if err != nil {
			global.LOG.Error("read body from request error:", zap.Error(err))
		} else {
			c.Request.SetBody(body)
		}
	} else {
		fmt.Println("GET")
		query := string(c.Request.URI().QueryString())
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		m := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		body, _ = json.Marshal(&m)
	}

	claims, _ := utils.GetClaims2(c)

	if claims.ID != 0 {
		userId = int(claims.ID)
	} else {
		id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		if err != nil {
			userId = 0
		}
		userId = id
	}
	record := system.SysOperationRecord{
		Ip:     c.ClientIP(),
		Method: string(c.Request.Method()),
		Path:   string(c.Request.URI().Path()),
		Agent:  string(c.UserAgent()),
		Body:   string(body),
		UserID: userId,
	}
	// 上传文件时候 中间件日志进行裁断操作
	if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
		if len(record.Body) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Body)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}

	writer := responseBodyWriter{
		ResponseWriter: c.GetWriter(),
		body:           &bytes.Buffer{},
	}
	now := time.Now()
	c.Next(ctx)
	latency := time.Since(now)
	record.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
	record.Status = c.GetResponse().StatusCode()
	record.Latency = latency
	record.Resp = writer.body.String()
	if strings.Contains(string(c.GetHeader("Pragma")), "public") ||
		strings.Contains(string(c.GetHeader("Expires")), "0") ||
		strings.Contains(string(c.GetHeader("Cache-Control")), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/force-download") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/octet-stream") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/vnd.ms-excel") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/download") ||
		strings.Contains(string(c.GetHeader("Content-Disposition")), "attachment") ||
		strings.Contains(string(c.GetHeader("Content-Transfer-Encoding")), "binary") {
		if len(record.Resp) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Resp)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}

	if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		global.LOG.Error("create operation record error:", zap.Error(err))
	}
}

func REC_VT(ctx context.Context, c *app.RequestContext) {
	fmt.Println("鉴权 AND 记录")
	claims, _ := utils.GetClaims2(c)

	// 验证身份信息x-token正确性
	var AuthorityId uint
	if claims.AuthorityId != 0 {
		AuthorityId = uint(int(claims.AuthorityId))
	} else {
		id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		if err != nil {
			AuthorityId = 0
		}
		AuthorityId = uint(id)
	}
	if claims.AuthorityId == 0 {
		return
	}
	var action common.PermissionType
	switch string(c.Request.Method()) {
	case http.MethodPost:
		action = common.ADD
		fmt.Println("新增操作类型")
	case http.MethodDelete:
		action = common.DELETE
		fmt.Println("删除操作类型")
	case http.MethodPut:
		action = common.UPDATE
		fmt.Println("修改操作类型")
	case http.MethodGet:
		action = common.GET
		fmt.Println("查询操作类型")
	default:
		fmt.Println("系统未定义操作类型")
		return
	}
	fmt.Println("路径:", string(c.Request.URI().Path()))
	reg, err := regexp.Compile(`/.*?/`)
	if err != nil {
		fmt.Println("正则表达式有误")
		return
	}
	menuName := c.Request.URI().Path()
	menuName = reg.Find(menuName)
	menuName = menuName[1 : len(menuName)-1]
	fmt.Println("正则匹配menuName：", string(menuName))
	err = system2.CheckAuth(action, AuthorityId, string(menuName))
	if err != nil {
		fmt.Println("权限验证未通过")
		return
	}
	c.Set("pass", true)
	// 3代表GET操作类型，该类型操作不进行记录
	if action == 3 {
		return
	}
	// 其他操作类型, 进行记录
	var body []byte
	var userId int
	if string(c.Request.Method()) != http.MethodGet {
		fmt.Println("非GET")
		fmt.Println("POST")
		var err error
		fmt.Println(c.Body())
		body, err = c.Body()
		if err != nil {
			global.LOG.Error("read body from request error:", zap.Error(err))
		} else {
			c.Request.SetBody(body)
		}
	} else {
		fmt.Println("GET")
		query := string(c.Request.URI().QueryString())
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		m := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		body, _ = json.Marshal(&m)
	}
	record := system.SysOperationRecord{
		Ip:     c.ClientIP(),
		Method: string(c.Request.Method()),
		Path:   string(c.Request.URI().Path()),
		Agent:  string(c.UserAgent()),
		Body:   string(body),
		UserID: userId,
	}
	// 上传文件时候 中间件日志进行裁断操作
	if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
		if len(record.Body) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Body)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	writer := responseBodyWriter{
		ResponseWriter: c.GetWriter(),
		body:           &bytes.Buffer{},
	}
	now := time.Now()
	c.Next(ctx)
	latency := time.Since(now)
	record.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
	record.Status = c.GetResponse().StatusCode()
	record.Latency = latency
	record.Resp = writer.body.String()
	if strings.Contains(string(c.GetHeader("Pragma")), "public") ||
		strings.Contains(string(c.GetHeader("Expires")), "0") ||
		strings.Contains(string(c.GetHeader("Cache-Control")), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/force-download") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/octet-stream") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/vnd.ms-excel") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/download") ||
		strings.Contains(string(c.GetHeader("Content-Disposition")), "attachment") ||
		strings.Contains(string(c.GetHeader("Content-Transfer-Encoding")), "binary") {
		if len(record.Resp) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Resp)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		global.LOG.Error("create operation record error:", zap.Error(err))
	}
}

func AuthAndRecordAssignHttp(ctx context.Context, c *app.RequestContext, permissionType common.PermissionType, apiName string) error {
	fmt.Println("鉴权 AND 记录")
	claims, err := utils.GetClaims2(c)
	if err != nil {
		return err
	}
	// 验证身份信息x-token正确性
	var AuthorityId uint
	if claims.AuthorityId != 0 {
		AuthorityId = uint(int(claims.AuthorityId))
	} else {
		id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		if err != nil {
			return err
		}
		AuthorityId = uint(id)
	}
	var action common.PermissionType
	switch string(c.Request.Method()) {
	case http.MethodPost:
		action = common.ADD
	case http.MethodDelete:
		action = common.DELETE
	case http.MethodPut:
		action = common.UPDATE
	case http.MethodGet:
		action = common.GET
	}
	err = system2.CheckAuth(action, AuthorityId, apiName)
	if err != nil {
		return err
	}
	// 代表GET操作类型，该类型操作不进行记录
	if action == 3 {
		return nil
	}
	// 其他操作类型, 进行记录
	var body []byte
	var userId int
	fmt.Println(c.Body())
	body, err = c.Body()
	if err != nil {
		global.LOG.Error("read body from request error:", zap.Error(err))
	} else {
		c.Request.SetBody(body)
	}
	record := system.SysOperationRecord{
		Ip:     c.ClientIP(),
		Method: string(c.Request.Method()),
		Path:   string(c.Request.URI().Path()),
		Agent:  string(c.UserAgent()),
		Body:   string(body),
		UserID: userId,
	}
	// 上传文件时候 中间件日志进行裁断操作
	if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
		if len(record.Body) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Body)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	writer := responseBodyWriter{
		ResponseWriter: c.GetWriter(),
		body:           &bytes.Buffer{},
	}
	now := time.Now()
	c.Next(ctx)
	latency := time.Since(now)
	record.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
	record.Status = c.GetResponse().StatusCode()
	record.Latency = latency
	record.Resp = writer.body.String()
	if strings.Contains(string(c.GetHeader("Pragma")), "public") ||
		strings.Contains(string(c.GetHeader("Expires")), "0") ||
		strings.Contains(string(c.GetHeader("Cache-Control")), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/force-download") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/octet-stream") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/vnd.ms-excel") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/download") ||
		strings.Contains(string(c.GetHeader("Content-Disposition")), "attachment") ||
		strings.Contains(string(c.GetHeader("Content-Transfer-Encoding")), "binary") {
		if len(record.Resp) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Resp)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		global.LOG.Error("create operation record error:", zap.Error(err))
	}
	return nil
}

func AuthAndRecord(ctx context.Context, c *app.RequestContext, permissionType common.PermissionType, apiName string) error {
	fmt.Println("鉴权 AND 记录")
	claims, err := utils.GetClaims2(c)
	if err != nil {
		return err
	}
	// 验证身份信息x-token正确性
	var AuthorityId uint
	if claims.AuthorityId != 0 {
		AuthorityId = uint(int(claims.AuthorityId))
	} else {
		id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		if err != nil {
			return err
		}
		AuthorityId = uint(id)
	}
	err = system2.CheckAuth(permissionType, AuthorityId, apiName)
	if err != nil {
		return err
	}
	// 代表GET操作类型，该类型操作不进行记录
	if string(c.Request.Method()) == http.MethodGet {
		return nil
	}
	// 其他操作类型, 进行记录
	var body []byte
	var userId int
	fmt.Println(c.Body())
	body, err = c.Body()
	if err != nil {
		global.LOG.Error("read body from request error:", zap.Error(err))
	} else {
		c.Request.SetBody(body)
	}
	record := system.SysOperationRecord{
		Ip:     c.ClientIP(),
		Method: string(c.Request.Method()),
		Path:   string(c.Request.URI().Path()),
		Agent:  string(c.UserAgent()),
		Body:   string(body),
		UserID: userId,
	}
	// 上传文件时候 中间件日志进行裁断操作
	if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
		if len(record.Body) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Body)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	writer := responseBodyWriter{
		ResponseWriter: c.GetWriter(),
		body:           &bytes.Buffer{},
	}
	now := time.Now()
	c.Next(ctx)
	latency := time.Since(now)
	record.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
	record.Status = c.GetResponse().StatusCode()
	record.Latency = latency
	record.Resp = writer.body.String()
	if strings.Contains(string(c.GetHeader("Pragma")), "public") ||
		strings.Contains(string(c.GetHeader("Expires")), "0") ||
		strings.Contains(string(c.GetHeader("Cache-Control")), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/force-download") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/octet-stream") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/vnd.ms-excel") ||
		strings.Contains(string(c.GetHeader("Content-Type")), "application/download") ||
		strings.Contains(string(c.GetHeader("Content-Disposition")), "attachment") ||
		strings.Contains(string(c.GetHeader("Content-Transfer-Encoding")), "binary") {
		if len(record.Resp) > 1024 {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, record.Resp)
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
	}
	if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		global.LOG.Error("create operation record error:", zap.Error(err))
	}
	return nil
}
