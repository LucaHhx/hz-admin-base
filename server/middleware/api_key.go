package middleware

import (
	"time"

	"hab/global"
	"hab/model/common/response"
	systemReq "hab/model/system/request"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ApiKeyOrJWT 组合认证中间件
// 优先检查 x-api-key 请求头，命中则注入管理员 claims 并跳过 JWT 认证
// 未命中则 fallback 到标准 JWTAuth 流程
func ApiKeyOrJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		configKey := global.HAB_CONFIG.AutoCode.ApiKey

		// 有有效 API Key 时，注入 claims 并跳过 JWT
		if configKey != "" && apiKey != "" {
			if apiKey != configKey {
				response.NoAuth("invalid api key", c)
				c.Abort()
				return
			}

			// 命中 API Key，注入超级管理员 claims
			claims := &systemReq.CustomClaims{
				BaseClaims: systemReq.BaseClaims{
					UUID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					ID:          1,
					Username:    "api-key",
					NickName:    "API Key User",
					AuthorityId: 1,
				},
				BufferTime: 0,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				},
			}
			c.Set("claims", claims)
			c.Next()
			return
		}

		// 无 API Key 或未配置，走标准 JWT 流程
		JWTAuth()(c)
	}
}
