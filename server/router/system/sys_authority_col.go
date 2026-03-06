package system

import (
	v1 "hab/api/v1"
	"hab/middleware"

	"github.com/gin-gonic/gin"
)

type AuthorityColRouter struct{}

var AuthorityColRouterApp = new(AuthorityColRouter)

func (s *AuthorityColRouter) InitAuthorityColRouter(privateGroup, publicGroup *gin.RouterGroup) {
	authorityRouter := privateGroup.Group("authorityCol").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := privateGroup.Group("authorityCol")
	var authorityColApi = v1.ApiGroupApp.SystemApiGroup.AuthorityColApi
	{
		authorityRouter.POST("setAuthorityCol", authorityColApi.SetAuthorityCol)             // 设置角色列权限
		authorityRouter.POST("canRemoveAuthorityCol", authorityColApi.CanRemoveAuthorityCol) // 删除角色列权限
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityCol", authorityColApi.GetAuthorityCol) // 获取角色列权限
		authorityRouterWithoutRecord.GET("getMenuColumns", authorityColApi.GetMenuColumns)    // 获取菜单对应的列
	}
}
