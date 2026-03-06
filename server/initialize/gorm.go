package initialize

import (
	"os"

	"hab/global"
	"hab/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.HAB_CONFIG.System.DbType {
	case "mysql":
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Pgsql.Dbname
		return GormPgSql()
	case "oracle":
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Oracle.Dbname
		return GormOracle()
	case "mssql":
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Mssql.Dbname
		return GormMssql()
	case "sqlite":
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Sqlite.Dbname
		return GormSqlite()
	default:
		global.HAB_ACTIVE_DBNAME = &global.HAB_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.HAB_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBindSession{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAuthorityCol{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},
		system.SysParams{},

		system.SysTableColumns{},

		system.SysDataFilter{},
	)
	if err != nil {
		global.HAB_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.HAB_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.HAB_LOG.Info("register table success")
}
