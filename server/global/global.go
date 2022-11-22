package global

import (
	"sync"

	"github.com/edufriendchen/hertz-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/edufriendchen/hertz-vue-admin/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB                      *gorm.DB
	DBList                  map[string]*gorm.DB
	REDIS                   *redis.Client
	CONFIG                  config.Server
	VIPER                   *viper.Viper
	LOG                     *zap.Logger
	TIMER                   timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
	EmailConfig             config.Email
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
