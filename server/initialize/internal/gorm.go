package internal

import (
	"hab/config"
	"hab/global"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch global.HAB_CONFIG.System.DbType {
	case "mysql":
		general = global.HAB_CONFIG.Mysql.GeneralDB
	case "pgsql":
		general = global.HAB_CONFIG.Pgsql.GeneralDB
	case "oracle":
		general = global.HAB_CONFIG.Oracle.GeneralDB
	case "sqlite":
		general = global.HAB_CONFIG.Sqlite.GeneralDB
	case "mssql":
		general = global.HAB_CONFIG.Mssql.GeneralDB
	default:
		general = global.HAB_CONFIG.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
