//go:build windows
// +build windows

package core

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	server2 "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

func initServer(opts ...config.Option) *server2.Hertz {
	h := server2.New(opts...)
	h.Use(recovery.Recovery())
	return h
}
