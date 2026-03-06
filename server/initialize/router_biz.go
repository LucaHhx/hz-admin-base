package initialize

import (
	"hz-admin-base/router"

	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		systemRouter := router.RouterGroupApp.System
		systemRouter.InitSysTableColumnsRouter(privateGroup, publicGroup)
		systemRouter.InitSysDataFilterRouter(privateGroup, publicGroup)
	}
}
