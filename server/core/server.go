package core

import (
	"fmt"
	"hab/global"
	"hab/initialize"
	"hab/service/system"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 从db加载jwt数据
	if global.HAB_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.HAB_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.HAB_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	HAB (Hz Admin Base) 后台管理系统
	Swagger 文档地址: http://127.0.0.1%s/swagger/index.html
	前端运行地址: http://127.0.0.1:8080
`, address)
	global.HAB_LOG.Error(s.ListenAndServe().Error())
}

// InitConf 仅初始化必要配置 不初始化数据库等服务 用于测试
func InitConf() {
	global.HAB_VP = Viper() // 初始化Viper
	initialize.OtherInit()
	global.HAB_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.HAB_LOG)

	global.InitSysMode()

	var err error
	if time.Local, err = time.LoadLocation(global.HAB_CONFIG.System.TimeZone); err != nil {
		panic(err)
	}
}

func InitServer() {
	InitConf()
	global.HAB_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.HAB_CONFIG.System.Migration {
		initialize.RegisterTables() // 初始化表
	}

	global.NoLogDB = global.HAB_DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Error)})

	if global.HAB_CONFIG.System.UseMultipoint || global.HAB_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		//initialize.RedisList()
	}

	if global.HAB_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
}

func RunApiServer() {
	// 从db加载jwt数据
	if global.HAB_DB != nil {
		system.LoadAll()
	}
	Router := initialize.ApiRouters()
	apiAddr := global.HAB_CONFIG.System.ApiAddr
	address := fmt.Sprintf(":%d", apiAddr)
	s := initServer(address, Router)
	global.HAB_LOG.Info("api server run success on ", zap.String("address", address))
	global.HAB_LOG.Error(s.ListenAndServe().Error())
}
