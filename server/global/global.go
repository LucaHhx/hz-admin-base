package global

import (
	"sync"

	"hab/utils/timer"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"hab/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	HAB_DB        *gorm.DB
	HAB_DBList    map[string]*gorm.DB
	NoLogDB       *gorm.DB
	HAB_REDIS     redis.UniversalClient
	HAB_REDISList map[string]redis.UniversalClient
	HAB_MONGO     *qmgo.QmgoClient
	HAB_CONFIG    config.Server
	HAB_VP        *viper.Viper
	// HAB_LOG    *oplogging.Logger
	HAB_LOG                 *zap.Logger
	HAB_Timer               timer.Timer = timer.NewTimerTask()
	HAB_Concurrency_Control             = &singleflight.Group{}
	HAB_ROUTERS             gin.RoutesInfo
	HAB_ACTIVE_DBNAME       *string
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
	Json                    = sonic.ConfigFastest
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return HAB_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := HAB_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func SkipLog(skip int) *zap.Logger {
	return HAB_LOG.WithOptions(zap.AddCallerSkip(skip))
}
