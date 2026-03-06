package core

import (
	"fmt"
	"hz-admin-base/global"
	"hz-admin-base/initialize"
	"hz-admin-base/service/system"
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
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-vue-admin
	当前版本:v2.7.9
	加群方式:微信号：shouzi_1994 QQ群：470239250
	项目地址：https://github.com/flipped-aurora/gin-vue-admin
	插件市场:https://plugin.gin-vue-admin.com
	GVA讨论社区:https://support.qq.com/products/371961
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	--------------------------------------版权声明--------------------------------------
	** 版权所有方：flipped-aurora开源团队 **
	** 版权持有公司：北京翻转极光科技有限责任公司 **
	** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

// InitConf 仅初始化必要配置 不初始化数据库等服务 用于测试
func InitConf() {
	global.GVA_VP = Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.InitSysMode()

	var err error
	if time.Local, err = time.LoadLocation(global.GVA_CONFIG.System.TimeZone); err != nil {
		panic(err)
	}
}

func InitServer() {
	InitConf()
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_CONFIG.System.Migration {
		initialize.RegisterTables() // 初始化表
	}

	global.NoLogDB = global.GVA_DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Error)})

	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		//initialize.RedisList()
	}

	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
}

func RunApiServer() {
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	Router := initialize.ApiRouters()
	apiAddr := global.GVA_CONFIG.System.ApiAddr
	address := fmt.Sprintf(":%d", apiAddr)
	s := initServer(address, Router)
	global.GVA_LOG.Info("api server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
