package middleware

import (
	"strconv"
	"strings"

	"hab/global"
	"hab/model/common/response"
	systemReq "hab/model/system/request"
	"hab/service"
	"hab/utils"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var waitUse *systemReq.CustomClaims
		if claims, exists := c.Get("claims"); exists {
			waitUse = claims.(*systemReq.CustomClaims)
		} else {
			waitUse, _ = utils.GetClaims(c)
		}
		if waitUse == nil {
			response.FailWithDetailed(gin.H{}, "insufficient_permissions", c)
			c.Abort()
			return
		}
		if waitUse.AuthorityId == 1 {
			// 如果为超级管理员则直接放行
			c.Next()
			return
		}
		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.HAB_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := casbinService.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			response.FailWithDetailed(gin.H{}, "insufficient_permissions", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
