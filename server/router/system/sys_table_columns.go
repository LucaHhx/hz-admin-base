package system

import (
	"hab/middleware"

	"github.com/gin-gonic/gin"
)

type SysTableColumnsRouter struct{}

// InitSysTableColumnsRouter 初始化 表字段配置 路由信息
func (s *SysTableColumnsRouter) InitSysTableColumnsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysTableColumnsRouter := Router.Group("sysTableColumns").Use(middleware.OperationRecord())
	sysTableColumnsRouterWithoutRecord := Router.Group("sysTableColumns")
	sysTableColumnsRouterWithoutAuth := PublicRouter.Group("sysTableColumns")
	{
		sysTableColumnsRouter.POST("createSysTableColumns", sysTableColumnsApi.CreateSysTableColumns)             // 新建表字段配置
		sysTableColumnsRouter.DELETE("deleteSysTableColumns", sysTableColumnsApi.DeleteSysTableColumns)           // 删除表字段配置
		sysTableColumnsRouter.DELETE("deleteSysTableColumnsByIds", sysTableColumnsApi.DeleteSysTableColumnsByIds) // 批量删除表字段配置
		sysTableColumnsRouter.PUT("updateSysTableColumns", sysTableColumnsApi.UpdateSysTableColumns)              // 更新表字段配置
		sysTableColumnsRouter.PUT("updateSysTableColumnsInfo", sysTableColumnsApi.UpdateSysTableColumnsInfo)      // 更新表字段配置
		sysTableColumnsRouter.GET("syncSysTableColumnsInfo", sysTableColumnsApi.SyncSysTableColumnsInfo)          // 更新表字段配置
		sysTableColumnsRouter.POST("importSysTableColumns", sysTableColumnsApi.ImportSysTableColumns)             // 导入表字段配置
	}
	{
		sysTableColumnsRouterWithoutRecord.GET("findSysTableColumns", sysTableColumnsApi.FindSysTableColumns)       // 根据ID获取表字段配置
		sysTableColumnsRouterWithoutRecord.GET("getSysTableColumnsList", sysTableColumnsApi.GetSysTableColumnsList) // 获取表字段配置列表
		sysTableColumnsRouterWithoutRecord.GET("getStructNameColumns", sysTableColumnsApi.GetStructNameColumns)     // 获取表字段配置列表

	}
	{
		sysTableColumnsRouterWithoutAuth.GET("getSysTableColumnsPublic", sysTableColumnsApi.GetSysTableColumnsPublic) // 表字段配置开放接口
	}
}
