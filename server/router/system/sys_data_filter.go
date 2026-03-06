package system

import (
	"hz-admin-base/middleware"
	"github.com/gin-gonic/gin"
)

type SysDataFilterRouter struct{}

// InitSysDataFilterRouter 初始化 sysDataFilter 路由信息
func (s *SysDataFilterRouter) InitSysDataFilterRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysDataFilterRouter := Router.Group("sysDataFilter").Use(middleware.OperationRecord())
	sysDataFilterRouterWithoutRecord := Router.Group("sysDataFilter")
	sysDataFilterRouterWithoutAuth := PublicRouter.Group("sysDataFilter")
	{
		sysDataFilterRouter.POST("createSysDataFilter", sysDataFilterApi.CreateSysDataFilter)             // 新建sysDataFilter
		sysDataFilterRouter.DELETE("deleteSysDataFilter", sysDataFilterApi.DeleteSysDataFilter)           // 删除sysDataFilter
		sysDataFilterRouter.DELETE("deleteSysDataFilterByIds", sysDataFilterApi.DeleteSysDataFilterByIds) // 批量删除sysDataFilter
		sysDataFilterRouter.PUT("updateSysDataFilter", sysDataFilterApi.UpdateSysDataFilter)              // 更新sysDataFilter
		sysDataFilterRouter.POST("importSysDataFilter", sysDataFilterApi.ImportSysDataFilter)             // 导入sysDataFilter
		sysDataFilterRouter.POST("filterData", sysDataFilterApi.FilterData)                               // 导出sysDataFilter
	}
	{
		sysDataFilterRouterWithoutRecord.GET("findSysDataFilter", sysDataFilterApi.FindSysDataFilter)        // 根据ID获取sysDataFilter
		sysDataFilterRouterWithoutRecord.POST("getSysDataFilterList", sysDataFilterApi.GetSysDataFilterList) // 获取sysDataFilter列表
		sysDataFilterRouterWithoutRecord.GET("executeSql", sysDataFilterApi.ExecuteSql)                      // 获取当前用户的sysDataFilter列表
	}
	{
		sysDataFilterRouterWithoutAuth.GET("getSysDataFilterPublic", sysDataFilterApi.GetSysDataFilterPublic) // sysDataFilter开放接口
	}
}
