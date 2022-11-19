package initialize

import (
	"context"
	"fmt"

	"github.com/edufriendchen/hertz-vue-admin/server/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err:", zap.Error(err))
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("redis connect ping response:", zap.String("pong", pong))
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
