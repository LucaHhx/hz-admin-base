package system

import (
	v1 "hab/api/v1"
	"hab/middleware"

	"github.com/gin-gonic/gin"
)

type AuthorityBtnRouter struct{}

var AuthorityBtnRouterApp = new(AuthorityBtnRouter)

func (s *AuthorityBtnRouter) InitAuthorityBtnRouter(privateGroup, publicGroup *gin.RouterGroup) {
	authorityRouter := privateGroup.Group("authorityBtn").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := privateGroup.Group("authorityBtn")
	var authorityBtnApi = v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouter.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouter.POST("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn)
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.POST("getAuthorityBtnAll", authorityBtnApi.GetAuthorityBtnAll)
	}
}
